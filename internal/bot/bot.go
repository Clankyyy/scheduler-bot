package bot

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	token string
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

	userNumbers := make(map[int64]int)

	bot.Handle("/start", func(c tele.Context) error {
		buttons := []telebot.Btn{
			tele.Btn{Text: "1", Data: "1"}, tele.Btn{Text: "2", Data: "2"}, tele.Btn{Text: "3", Data: "3"},
		}
		menu := &tele.ReplyMarkup{}
		menu.Inline(
			menu.Row(buttons...),
		)
		// btn1 := menu.Text("1")
		// btn2 := menu.Text("2")
		return c.Send("Добро пожаловать выберите группу", menu)
	})

	// bot.Handle(&btnHelp, func(c tele.Context) error {
	// 	return c.Send("Got it")
	// })

	// On inline button pressed (callback)
	bot.Handle(telebot.OnCallback, func(c telebot.Context) error {
		// Get the callback data and parse it as an integer.
		fmt.Println("in callback")
		callbackData := c.Callback().Data
		number, err := strconv.Atoi(callbackData)
		if err != nil {
			return c.Send("Invalid number")
		}

		// Store the user's selected number.
		userNumbers[c.Sender().ID] = number

		// Send a confirmation message to the user.
		return c.Send(fmt.Sprintf("You selected the number %d", number))
	})

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		// Get the user's selected number from the map.
		number, ok := userNumbers[c.Sender().ID]
		if !ok {
			return c.Send("Please select a number first")
		}

		// Reply to the user with the number they selected.
		return c.Send(fmt.Sprintf("Your selected number is %d", number))
	})

	bot.Start()
}

func (b *Bot) handleStartLogin(c tele.Context) error {
	buttons := []telebot.Btn{
		tele.Btn{Text: "1", Data: "1"}, tele.Btn{Text: "2", Data: "2"}, tele.Btn{Text: "3", Data: "3"},
	}
	menu := &tele.ReplyMarkup{}
	menu.Inline(
		menu.Row(buttons...),
	)
	// btn1 := menu.Text("1")
	// btn2 := menu.Text("2")
	return c.Send("Добро пожаловать выберите группу", menu)
}
