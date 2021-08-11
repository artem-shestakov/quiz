package model

type Answer struct {
	ID int `json:"id"`
	QuestionID int `json:"question_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Content string `json:"content"`
	IsCorrect bool `json:"is_correct"`
}
