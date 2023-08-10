// handlers/callbacks.go

package handlers

import (
	"Telegram-Store/services"
	"Telegram-Store/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, data := utils.GetKeyValue(update.CallbackQuery.Data)

	datauint, err := strconv.Atoi(data)

	if err != nil {
		// Обработка ошибки, если не удалось преобразовать в uint
		// Например, вы можете просто проигнорировать такой callback-запрос
		return
	}
	switch {
	case cmd == "borrow_book" || cmd == "return_book":
		services.BorrowBookCallback(bot, update, datauint)
	default:
		//TODO: rework logic
	}
}
