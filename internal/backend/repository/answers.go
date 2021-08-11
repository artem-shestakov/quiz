package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/artem-shestakov/quiz/internal/model"
)

type AnswerRepository struct {
	db *sqlx.DB
}

func NewAnswerRepository(db *sqlx.DB) *AnswerRepository {
	return &AnswerRepository{
		db: db,
	}
}

func (q *AnswerRepository) CreateAnswer(questionID int, content string, isCorrect bool)  (*model.Answer, error){
	var answer model.Answer
	query := fmt.Sprintf("INSERT INTO %s (question_id, content, is_correct) VALUES ($1, $2, $3) RETURNING *", answersTable)
	err := q.db.QueryRow(query, questionID, content, isCorrect).Scan(&answer.ID, &answer.QuestionID, &answer.CreatedAt, &answer.UpdatedAt, &answer.Content, &answer.IsCorrect)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}
