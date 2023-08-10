// repositories/books.go

package repositories

import (
	"Telegram-Store/errors"
	"Telegram-Store/models"
)

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if result := DB.Find(&books); result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func BorrowBook(bookID int, userID int64) error {
	var book models.Book
	if result := DB.First(&book, bookID); result.Error != nil {
		return result.Error
	}

	if !book.IsAvailable() {
		return errors.ErrBookNotAvailable
	}

	book.BorrowedBy = int(userID)
	if result := DB.Save(&book); result.Error != nil {
		return result.Error
	}

	return nil
}

func ReturnBook(bookID int) error {
	var book models.Book
	if result := DB.First(&book, bookID); result.Error != nil {
		return result.Error
	}

	book.BorrowedBy = 0
	if result := DB.Save(&book); result.Error != nil {
		return result.Error
	}

	return nil
}

func GetBookByID(bookID int) (*models.Book, error) {
	var book models.Book
	if result := DB.First(&book, bookID); result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}
