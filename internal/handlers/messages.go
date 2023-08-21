package handlers

import (
	"Telegram-Store/internal/services"
	"Telegram-Store/internal/storage/repositories"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	services.HandleMessage(bot, update)
}

func MessagesInput(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, userRepo repositories.User, bookRepo repositories.Book) {
	fmt.Println("Inside MessagesInput")
	services.HandleUserTakeBook(ctx, bot, update, userRepo, bookRepo)
}
