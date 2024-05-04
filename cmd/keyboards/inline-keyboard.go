package keyboards

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

func ActivateInlineKeyboard() {
	keyboard := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("Start to learn Go!").
				WithRequestContact(true).WithWebApp("https://go.dev"),
			tu.KeyboardButton("Start to learn Python!").
				WithRequestContact(true).WithWebApp("https://python.org"),
		),
	)
}
