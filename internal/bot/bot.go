package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/Clankyyy/scheduler-bot/internal/current"
	"github.com/Clankyyy/scheduler-bot/internal/markup"
	"github.com/Clankyyy/scheduler-bot/internal/schedule"
	"github.com/Clankyyy/scheduler-bot/internal/storage"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	token   string
	current current.Currenter
	storage storage.Storager
}

func NewBot(token string, current current.Currenter, storage storage.Storager) *Bot {
	return &Bot{
		token:   token,
		current: current,
		storage: storage,
	}
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
	btnDaily := mainMenu.Text("ℹ Расписание на день")
	btnWeekly := mainMenu.Text("✅ Расписание на неделю")
	mainMenu.Reply(
		mainMenu.Row(btnDaily),
		mainMenu.Row(btnWeekly),
	)

	bot.Handle("/start", b.handleStart)

	// Handles group select
	bot.Handle(tele.OnCallback, func(c tele.Context) error {
		selectedGroup := c.Callback().Data
		err := b.storage.AddUser(c.Sender().ID, selectedGroup)
		if err != nil {
			log.Println("Error during AddUser", err.Error())
			return c.Send("Произошла ошибка при регистрации, пожалуйста попробуйте позже")
		}
		return c.Send(fmt.Sprintf("Вы выбрали группу %s", selectedGroup), mainMenu)
	})

	bot.Handle(&btnDaily, b.handleGetDaily)
	bot.Handle(&btnWeekly, b.handleGetWeekly)

	bot.Start()
}

func (b *Bot) handleGetDaily(c tele.Context) error {
	day, weekType := b.current.Now()
	slug, err := b.storage.GetSlug(c.Sender().ID)
	if err != nil {
		log.Println("Cant read slug from db:", err.Error())
		return c.Send("Расписание для вас в данный момент недоступно, вероятно вы не выбрали группу")
	}
	daily, err := schedule.GetDaily(slug, day, weekType)
	if err != nil {
		log.Println("Cant get daily from backend: ", err.Error())
		return c.Send("Произошла ошибка, попробуйте позже")
	}

	// log.Print(daily)

	return c.Send(&daily)
}

func (b *Bot) handleGetWeekly(c tele.Context) error {
	_, weekType := b.current.Now()
	slug, err := b.storage.GetSlug(c.Sender().ID)
	if err != nil {
		log.Println("Cant read slug from db:", err.Error())
		return c.Send("Расписание для вас в данный момент недоступно, вероятно вы не выбрали группу")
	}
	daily, err := schedule.GetWeekly(slug, weekType)
	if err != nil {
		log.Println("Cant get daily from backend: ", err.Error())
		return c.Send("произошла ошибка, попробуйте позже")
	}

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
