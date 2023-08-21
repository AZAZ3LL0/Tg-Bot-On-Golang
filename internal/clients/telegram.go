package clients

import (
	"github.com/spf13/viper"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tgbotapi.BotAPI {

	apiToken := viper.GetString("bot.token")

	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Println(err)
		return nil
	}

	bot.Debug = true

	return bot
}
