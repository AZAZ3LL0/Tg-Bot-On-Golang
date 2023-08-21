package repositories

import (
	"context"
	"database/sql"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserTgId(ctx context.Context, update tgbotapi.Update) error {
	tgUserID := update.Message.From.ID

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE tg_id = $1", tgUserID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	username := update.Message.From.UserName
	result, err := r.db.ExecContext(ctx, "INSERT INTO users (name, tg_id) VALUES ($1, $2)", username, tgUserID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows != 1 {
		return err
	}
	return nil
}

func (r *userRepository) BorrowBook(ctx context.Context, bookId int, userId int64) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var existingBookID int
	err = tx.GetContext(ctx, &existingBookID, `
	   SELECT book_id
	   FROM user_book
	   WHERE user_id = $1 AND book_id = $2`, userId, bookId)
	if err == nil {
		return errors.New("user already borrowed this book")
	} else if err != sql.ErrNoRows {
		return err
	}

	_, err = tx.ExecContext(ctx, `
        UPDATE book
        SET quantity = quantity - 1
        WHERE id = $1 AND quantity > 0`, bookId)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
        INSERT INTO user_book (user_id, book_id)
        VALUES ($1, $2)`, userId, bookId)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
