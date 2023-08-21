package services

import (
	"Telegram-Store/internal/storage/repositories"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func ShowAllBooks(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, repo repositories.Book) {
	fmt.Println("Inside ShowAllBooks function")
	books, err := repo.GetAllBooks(ctx)
	if err != nil {
		log.Println("Error retrieving books:", err)
		text := "An error occurred while processing your request. Please try again later."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		_, _ = bot.Send(msg)
		return
	}

	if len(books) == 0 {
		text := "There are no available books at the moment."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		_, _ = bot.Send(msg)
		return
	}

	var bookListBuilder strings.Builder
	bookListBuilder.WriteString("Available books:\n")
	for _, book := range books {
		bookListBuilder.WriteString(fmt.Sprintf("%d. %s - %s\n", book.ID, book.Title, book.Author))
	}

	text := bookListBuilder.String()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	_, _ = bot.Send(msg)
}

func GetAuthorsBooks(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, repo repositories.Book) {

}

func BorrowBookCallback(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, datalist string) {
}
