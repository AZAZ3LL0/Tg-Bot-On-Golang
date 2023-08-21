package models

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	TgId  string `json:"tg_id"`
	Books []Book `json:"books"`
}
