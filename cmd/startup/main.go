package main

import (
	"fmt"
	"log"
	config "telegram-bot"
	"telegram-bot/cmd/sticker"

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

	handler.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, _ = bot.SendMessage(
			tu.Message(
				tu.ID(message.Chat.ID),
				fmt.Sprintf("You said: %s", message.Text),
			),
		)
	}, th.AnyMessageWithText())

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Hello, %s!", update.Message.From.FirstName),
			),
		)
	}, th.CommandEqual("start"))

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendSticker(
			tu.Sticker(
				tu.ID(update.Message.Chat.ID),
				*sticker.GetRandomSticker(),
			),
		)
	}, th.CommandEqual("sticker"))

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
				"Unknown command, try /help to get list of commands.",
			),
		)
	}, th.AnyCommand())

	handler.Start()
}
