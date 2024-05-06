package sticker

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/mymmrac/telego"
)

var (
	stickers []string = []string{
		"CAACAgIAAxkBAAErOllmNx1x-cbZUzLqnxz9b2k1e4ICjQACmyMAAhGYAAFKrV6CanU7BCo1BA",
		"CAACAgIAAxkBAAErOldmNx0WQhJKm60b2_qYur-MgAb7mAACDUgAAqtDgUjYgUC3Y2n5VTUE",
		"CAACAgIAAxkBAAErOlVmNx0ETaYzU9mdws8nggABi6T384kAAlE4AAKGnfBIFMIjLw9QOhM1BA",
		"CAACAgIAAxkBAAErOlNmNxzsQaTuOeiAKibSlhgdoRnusAACRg4AAjw-qEj0TUSy5rnZ2TUE",
		"CAACAgIAAxkBAAErOlFmNxzKFlqgSmCkyj9EKm1BBT1CMgACxAoAAo_9EUg92yX9mfyMPjUE",
		"CAACAgEAAxkBAAErOk9mNxyodnSlxcr8migaiuaKNNgIgAACDAIAAlm6eQx3Auohch7KRjUE",
		"CAACAgIAAxkBAAErGnhmLfgeJ5AMpNBHvYYIA63NGcNEDQACxAwAAitNsUji1ATkPfeWQTQE",
		"CAACAgIAAxkBAAErOmlmNyiBysyTDqSmfpGwNzdYHghw7QAC5CUAAsfyAAFKL2MmNArDN0M1BA",
		"CAACAgIAAxkBAAErOmtmNyiWB955-4tWMAxSDjKPmhyubQACwCoAAgxwAUrQCnAmGrcLFjUE",
	}
)

func GetRandomSticker() *telego.InputFile {
	max := big.NewInt(int64(len(stickers) - 1))
	randomNumber, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return &telego.InputFile{
		FileID: stickers[randomNumber.Int64()],
	}
}
