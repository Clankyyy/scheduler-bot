package bot

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Clankyyy/scheduler-bot/internal/markup"
	"github.com/Clankyyy/scheduler-bot/internal/schedule"
	tele "gopkg.in/telebot.v3"
)

var userNumbers map[int64]string

type Bot struct {
	token string
}

func init() {
	userNumbers = make(map[int64]string, 15)
}

func NewBot(token string) *Bot {
	return &Bot{token: token}
}

func (b *Bot) Start() {
	pref := tele.Settings{
		Token:  b.token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	// menus, buttons, etc..
	mainMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnHelp := mainMenu.Text("ℹ Help")
	btnSettings := mainMenu.Text("⚙ Settings")
	mainMenu.Reply(
		mainMenu.Row(btnHelp),
		mainMenu.Row(btnSettings),
	)

	bot.Handle("/start", b.handleStart)

	// Handles group select
	bot.Handle(tele.OnCallback, func(c tele.Context) error {

		selectedGroup := c.Callback().Data
		userNumbers[c.Sender().ID] = selectedGroup

		// Send a confirmation message to the user.
		return c.Send(fmt.Sprintf("Вы выбрали группу %s", selectedGroup), mainMenu)
	})

	bot.Handle(&btnHelp, b.handleGetDaily)
	bot.Handle(tele.OnText, b.handleGetSchedule)

	bot.Start()
}

func (b *Bot) handleGetSchedule(c tele.Context) error {
	group, ok := userNumbers[c.Sender().ID]
	if !ok {
		return c.Send("Пожалуйста, сначала выберете группу")
	}

	return c.Send(fmt.Sprintf("Ваша группа: %s", group))
}

func (b *Bot) handleGetDaily(c tele.Context) error {
	daily, err := schedule.GetDaily("2-4306", "monday", "even")
	if err != nil {
		log.Print(err.Error())
		return c.Send("АШИБКА")
	}

	// log.Print(daily)

	return c.Send(&daily)
}

func (b *Bot) handleStart(c tele.Context) error {
	groups, err := schedule.GetGroups()
	if err != nil {
		log.Print(err)
		return c.Send("Сервис недоступен, пожалуйста попробуйте позже")
	}
	groupButtons := markup.GroupList(groups)
	return c.Send("Выбирете группу", groupButtons)
}

type Test struct {
	num int
}

func (t *Test) Send(b *tele.Bot, r tele.Recipient, opts *tele.SendOptions) (*tele.Message, error) {
	str := strconv.Itoa(t.num)
	return b.Send(r, str)
}
