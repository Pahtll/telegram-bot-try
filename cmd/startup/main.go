package main

import (
	"fmt"
	"log"
	config "telegram-bot"
	"telegram-bot/cmd/handlers"
	"telegram-bot/cmd/keyboards"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	bot, err := telego.NewBot(config.API_KEY, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatalf("failed to create new bot: %v", err)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	handler, _ := th.NewBotHandler(bot, updates)
	defer handler.Stop()

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Hello, %s!", update.Message.From.FirstName),
			),
		)
	}, th.CommandEqual("start"))

	handler.Handle(handlers.SendRandSticker, th.CommandEqual("sticker"))

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Try /start to get your greet, /sticker to get random sticker or /help to get list of commands.\nMore commands will be added later",
			),
		)
	}, th.CommandEqual("help"))

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Reply keyboard, use commands if it's needed",
			).WithReplyMarkup(keyboards.ReplyKeyboard),
		)
	}, th.CommandEqual("reply"))

	handler.Handle(handlers.SendRandEmoji, th.CommandEqual("emoji"))

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Unknown command, try /help to get list of commands.",
			),
		)
	}, th.AnyCommand())

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		var max int64 = 3

		randomNumber := handlers.RandomNumber(max)

		switch randomNumber {
		case 0:
			handlers.SendRandImg(bot, update)
		case 1:
			handlers.SendRandSticker(bot, update)
		case 2:
			handlers.SendRandEmoji(bot, update)
		default:
			handlers.EchoMessage(bot, update)
		}
	}, th.AnyMessageWithText())

	handler.Start()
}
