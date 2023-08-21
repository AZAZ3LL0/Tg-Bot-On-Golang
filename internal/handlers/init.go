package handlers

import (
	"Telegram-Store/internal/storage/repositories"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(ctx context.Context, bot *tgbotapi.BotAPI, bookRepo repositories.Book, userRepo repositories.User) {
	fmt.Println("Inside Init function")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Println("Received callback query")
			Callbacks(ctx, bot, update)
		} else if update.Message.IsCommand() {
			fmt.Println("Received command")
			Commands(ctx, bot, update, bookRepo, userRepo)
		} else {
			fmt.Println("Received message")
			MessagesInput(ctx, bot, update, userRepo, bookRepo)
		}
	}
}
