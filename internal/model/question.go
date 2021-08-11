package model

type Question struct {
	ID int `json:"id"`
	Type string `json:"type"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	Content string `json:"content" db:"content"`
}
