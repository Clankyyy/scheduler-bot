package bot

import (
	"time"

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

	bot.Handle("/hi", func(c tele.Context) error {
		c.Respond()
		return c.Send("hello!")
	})

	bot.Start()
}
