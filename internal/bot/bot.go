package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

var userNumbers map[int64]int
var httpClient *http.Client

type Bot struct {
	token string
}

func init() {
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	userNumbers = make(map[int64]int, 15)
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

	bot.Handle("/start", b.handleStartLogic)

	// Handles group select
	bot.Handle(tele.OnCallback, func(c tele.Context) error {
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

	bot.Handle(tele.OnText, func(c tele.Context) error {
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

func (b *Bot) handleStartLogic(c tele.Context) error {
	type wrapper struct {
		result []GroupDataReq
		err    error
	}
	ch := make(chan wrapper, 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		result, err := requestGroupsWithContext(ctx)
		ch <- wrapper{result, err}
	}()

	select {
	case data := <-ch:
		str := data.result[0].Name + "-" + data.result[0].Course
		buttons := []tele.Btn{
			{Text: str, Data: "1"}, {Text: "2", Data: "2"}, {Text: "3", Data: "3"},
		}
		menu := &tele.ReplyMarkup{}
		menu.Inline(
			menu.Row(buttons...),
		)
		return c.Send("Добро пожаловать выберите группу", menu)
	case <-ctx.Done():
		return c.Send("Что то пошло не так")
	}
}

func requestGroupsWithContext(ctx context.Context) ([]GroupDataReq, error) {
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, "http://localhost:8000/schedule/", nil)
	if err != nil {
		panic(err)
	}

	res, err := httpClient.Do(req)
	var data []GroupDataReq
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

type GroupDataReq struct {
	Course string `json:"course"`
	Name   string `json:"name"`
}
