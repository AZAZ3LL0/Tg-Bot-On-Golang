package services

import (
	"Telegram-Store/internal/storage/repositories"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func TakeBook(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, userRepo repositories.User, bookRepo repositories.Book) {
	fmt.Println("Inside TakeBook function")

	txt := "Please enter the ID of the book you wish to borrow:"
	message := tgbotapi.NewMessage(update.Message.Chat.ID, txt)
	_, err := bot.Send(message)
	if err != nil {
		return
	}
	bookID, err := strconv.Atoi(update.Message.Text)

	err = userRepo.BorrowBook(ctx, bookID, update.Message.From.ID)
	if err != nil {
		log.Println("Error borrowing book:", err)
		text := "An error occurred while borrowing the book. Please try again later."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
		return
	}

	book, err := bookRepo.GetBookByID(ctx, bookID)
	if err != nil {
		log.Println("Error retrieving book:", err)
		text := "An error occurred while retrieving the book information. Please try again later."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
		return
	}

	text := "You've successfully picked up the book:\n" +
		"Title: " + book.Title + "\n" +
		"Author: " + book.Author + "\n" +
		"Genre: " + book.Genre + "\n" +
		"Description: " + book.Description
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	bot.Send(msg)
}

func ReturnBook(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("Inside ReturnBook function")
	text := "Please enter the ID of the book you wish to return."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}
