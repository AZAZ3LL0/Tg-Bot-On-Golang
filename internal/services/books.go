package services

import (
	"Telegram-Store/storage/repositories"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Привет! Добро пожаловать в онлайн библиотеку. Чтобы посмотреть список книг, используйте команду /show_books.\""
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Извините, я не понимаю эту команду. Попробуйте /show_books для просмотра доступных книг."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func ShowAllBooks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	books, err := repositories.GetAllBooks()
	if err != nil {
		text := "Произошла ошибка при получении списка книг."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		return
	}

	if len(books) == 0 {
		text := "В библиотеке пока нет доступных книг."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		return
	}

	text := "Список доступных книг:\n"
	for _, book := range books {
		availability := "Доступна"
		if !book.IsAvailable() {
			availability = "Недоступна"
		}
		text += fmt.Sprintf("ID: %d\nНазвание: %s\nАвтор: %s\nДоступность: %s\n\n", book.ID, book.Title, book.Author, availability)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func BorrowBook(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Введите ID книги, которую хотите взять в аренду:"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func ReturnBook(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Введите ID книги, которую хотите вернуть:"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func BorrowBookCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, bookId int) {
	userId := update.CallbackQuery.Message.Chat.ID

	// Проверяем доступность книги
	book, err := repositories.GetBookByID(bookId)
	if err != nil {
		// Обработка ошибки
		return
	}

	if book.Available == false {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "The book is not available for borrowing")
		if _, err := bot.Send(msg); err != nil {
			// Обработка ошибки отправки сообщения
			return
		}
		return
	}

	// Выдаем книгу пользователю
	err = repositories.BorrowBook(bookId, userId)
	if err != nil {
		// Обработка ошибки
		return
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("You have successfully borrowed the book: %s", book.Title))
	if _, err := bot.Send(msg); err != nil {
		// Обработка ошибки отправки сообщения
		return
	}
}

func ReturnBookCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, bookId int) {
	err := repositories.ReturnBook(bookId)
	if err != nil {
		// Обработка ошибки
		return
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Book successfully returned")
	if _, err := bot.Send(msg); err != nil {
		// Обработка ошибки отправки сообщения
		return
	}
}
