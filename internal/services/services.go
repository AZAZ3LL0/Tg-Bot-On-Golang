package services

import (
	"Telegram-Store/internal/keyboards"
	"Telegram-Store/internal/storage/repositories"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func Start(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, userRepo repositories.User) {
	err := userRepo.GetUserTgId(ctx, update)
	if err != nil {
		fmt.Println("Error getting user tgID:", err)
		return
	}
	text := "Hi, welcome to the online library. To see the list of commands, use the command: /help"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func Help(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Here all commands \n/take\n/return\n/show_all_books\n/show_my_books\n/show_books_by_author\n"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}

}

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Sorry, I don't understand this command. Try /help to view the available commands."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func HandleUserTakeBook(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, userRepo repositories.User, bookRepo repositories.Book) {
	fmt.Println("Inside HandleUserInput function")
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
