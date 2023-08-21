package repositories

import (
	"Telegram-Store/internal/storage/models"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type bookRepository struct {
	db *sqlx.DB
}

func NewbookRepository(db *sqlx.DB) *bookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	fmt.Println("Inside GetAllBooks function")
	var booksDTO []*models.Book

	rows, err := r.db.QueryxContext(ctx, `
		SELECT
  		book.id,
  		book.title,
  		book.author
		FROM book`)

	if err != nil {
		return nil, fmt.Errorf("%s", errParsRows)
	}

	defer rows.Close()

	for rows.Next() {
		bookDTO := &models.Book{}
		if err = rows.Scan(
			&bookDTO.ID,
			&bookDTO.Title,
			&bookDTO.Author,
		); err != nil {
			return nil, fmt.Errorf("%s", errScanRows)
		}

		booksDTO = append(booksDTO, bookDTO)
	}

	return booksDTO, nil
}

func (r *bookRepository) GetBooksByAuthor(ctx context.Context, author string) ([]*models.Book, error) {
	fmt.Println("Inside GetBookByAuthor function")
	var booksDTO []*models.Book

	query := `
		SELECT
		book.title,
		book.description,
		book.genre
		FROM book 
		WHERE book.author = $1`

	rows, err := r.db.QueryxContext(ctx, query, author)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		bookDTO := &models.Book{}
		if err = rows.Scan(
			&bookDTO.Title,
			&bookDTO.Description,
			&bookDTO.Genre,
		); err != nil {
			return nil, fmt.Errorf("%s", errScanRows)
		}

		booksDTO = append(booksDTO, bookDTO)
	}

	return booksDTO, nil
}

func (r *bookRepository) GetBookByID(ctx context.Context, bookID int) (*models.Book, error) {
	bookDTO := &models.Book{}

	row := r.db.QueryRowxContext(ctx, fmt.Sprintf(`
		SELECT
    	book.id,
    	book.title,
    	book.description,
    	book.author,
    	book.genre
    	FROM book
    	WHERE book.id = %d LIMIT 1`, bookID))

	if err := row.Scan(
		&bookDTO.ID,
		&bookDTO.Title,
		&bookDTO.Description,
		&bookDTO.Author,
		&bookDTO.Genre,
	); err != nil {
		return nil, fmt.Errorf("%s", errScanRow)
	}

	return bookDTO, nil
}
