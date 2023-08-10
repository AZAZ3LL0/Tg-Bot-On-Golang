package clients

import (
	"Telegram-Store/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tgbotapi.BotAPI {

	apiToke, err := config.Config("TELEGRAM_APITOKEN")

	if err != nil {
		log.Println(err)
		return nil
	}

	bot, err := tgbotapi.NewBotAPI(apiToke)
	if err != nil {
		log.Println(err)
		return nil
	}

	bot.Debug = true

	return bot
}
