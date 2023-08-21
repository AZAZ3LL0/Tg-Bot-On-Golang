package models

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Quantity    int    `json:"quantity"`
	BorrowerID  int64  `json:"borrower_id"`
}
