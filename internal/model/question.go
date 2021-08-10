package model

type Question struct {
	ID int `json:"id"`
	Type string `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Context string `json:"context"`
}
