package handlers

import (
	"Telegram-Store/internal/services"
	"Telegram-Store/internal/storage/repositories"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, bookRepo repositories.Book, userRepo repositories.User) {
	switch update.Message.Command() {
	case "help":
		services.Help(bot, update)
	case "start":
		fmt.Println("Received start command")
		services.Start(ctx, bot, update, userRepo)
	case "show_all_books":
		fmt.Println("Received show_all_books command")
		services.ShowAllBooks(ctx, bot, update, bookRepo)
	case "show_books_by_author":
		fmt.Println("Received show_books_by_author command")
		services.GetAuthorsBooks(ctx, bot, update, bookRepo)
	case "take":
		fmt.Println("Received take command")
		services.TakeBook(ctx, bot, update, userRepo, bookRepo)
	case "return":
		fmt.Println("Received return command")
		services.ReturnBook(bot, update)
	default:
		fmt.Println("Unknown command")
		Messages(bot, update)
	}
}
