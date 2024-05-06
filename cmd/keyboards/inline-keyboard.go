package keyboards

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

var ReplyKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("Start to learn Go!"),

		tu.KeyboardButton("Start to learn Python!").
			WithText("== Отсосать член..."),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("LOL"),
	),
)

var InlineKeyboard = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("What about C#?"),

		tu.InlineKeyboardButton("LOL").WithCallbackData("LOL"),
	),
)
