package handlers

import (
	"Telegram-Store/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "show_books":
		services.ShowAllBooks(bot, update)
	case "borrow_book":
		services.BorrowBook(bot, update)
	case "return_book":
		services.ReturnBook(bot, update)
	}
}
