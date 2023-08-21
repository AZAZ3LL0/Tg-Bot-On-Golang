package repositories

import (
	"Telegram-Store/internal/storage/models"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
)

const (
	errScanRows     = "error occurred scanning rows"
	errParsRows     = "error occurred parsing rows"
	errScanRow      = "error occurred scanning single row"
	errUserTookBook = "error occurred db: user already have this book"
)

type Repository struct {
	Book
	User
}

type Option func(r *Repository)

type Book interface {
	GetAllBooks(ctx context.Context) ([]*models.Book, error)
	GetBooksByAuthor(ctx context.Context, author string) ([]*models.Book, error)
	GetBookByID(ctx context.Context, bookID int) (*models.Book, error)
}

type User interface {
	BorrowBook(ctx context.Context, bookId int, userId int64) error
	GetUserTgId(ctx context.Context, update tgbotapi.Update) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Book: NewbookRepository(db),
		User: NewUserRepository(db),
	}
}
