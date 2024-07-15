package markup

import (
	"github.com/Clankyyy/scheduler-bot/internal/entity"
	tele "gopkg.in/telebot.v3"
)

func GroupList(groups []entity.GroupDataReq) *tele.ReplyMarkup {
	buttons := make([]tele.Btn, 0, len(groups))
	for _, g := range groups {
		data := g.Course + "-" + g.Name
		buttons = append(buttons, tele.Btn{Text: data, Data: data})
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
