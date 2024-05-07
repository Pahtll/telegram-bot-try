package handlers

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"

	"fmt"
	"telegram-bot/cmd/sticker"
)

func EchoMessage(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(
		tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("You said: %s", update.Message.Text),
		),
	)
}

func SendRandImg(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendPhoto(
		tu.Photo(
			tu.ID(update.Message.Chat.ID),
			*sticker.GetRandomPicture(),
		),
	)
}
func SendRandSticker(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendSticker(
		tu.Sticker(
			tu.ID(update.Message.Chat.ID),
			*sticker.GetRandomSticker(),
		),
	)
}

func SendRandEmoji(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(
		tu.Message(
			tu.ID(update.Message.Chat.ID),
			sticker.GetRandomEmoji(),
		),
	)
}

func RandomNumber(max int64) int64 {
	num := big.NewInt(max)

	randN, err := rand.Int(rand.Reader, num)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return randN.Int64()
}
