package main

import (
	"fmt"
	"log"
	"telegram-bot/cmd/sticker"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	bot, err := telego.NewBot("6701730667:AAGgiU_c_wmKAxonukecApZ8FVE2cPepkvs", telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatalf("failed to create new bot: %v", err)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if msg := update.Message; msg != nil {
			chatID := msg.Chat.ID
			if msg.Text == "/sticker" {
				sticker := tu.Sticker(msg.Chat.ChatID(), *sticker.GetRandomSticker())
				bot.SendSticker(sticker)
				fmt.Printf("Sticker sent to %v\n", msg.Chat.Username)
			} else {
				bot.SendMessage(tu.Message(tu.ID(chatID), msg.Text))
				fmt.Printf("Message sent to %v\n", msg.Chat.Username)
			}
		}
	}
}
