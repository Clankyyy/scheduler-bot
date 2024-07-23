package markup

import (
	"github.com/Clankyyy/scheduler-bot/internal/entity"
	tele "gopkg.in/telebot.v3"
)

func GroupList(groups entity.GroupsRes) *tele.ReplyMarkup {
	buttons := make([]tele.Btn, 0, len(groups.Data))
	for _, g := range groups.Data {
		buttons = append(buttons, tele.Btn{Text: g, Data: g})
	}
	menu := &tele.ReplyMarkup{}
	menu.Inline(menu.Split(3, buttons)...)
	return menu
}

func MainMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	// Reply buttons.
	btnHelp := menu.Text("ℹ Help")
	btnSettings := menu.Text("⚙ Settings")
	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)
	return menu
}
