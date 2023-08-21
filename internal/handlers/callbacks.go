package handlers

import (
	"Telegram-Store/internal/services"
	"Telegram-Store/internal/utils"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, data := utils.GetKeyValue(update.CallbackQuery.Data)
	switch {
	case cmd == "borrow_book" || cmd == "return_book":
		services.BorrowBookCallback(ctx, bot, update, data)
	default:
		//TODO: rework logic
	}
}
