package models

type Book struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Author      string `gorm:"not null"`
	Available   bool
	BorrowedBy  int
}

func (b *Book) IsAvailable() bool {
	return b.Available // Просто возвращаем значение поля Available
}

func (b *Book) BorrowedByID() int {
	return b.BorrowedBy // Просто возвращаем значение поля BorrowedBy
}
