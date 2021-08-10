package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const questionsTable = "questions"

type QuestionsRepository struct {
	db *sqlx.DB
}

func NewQuestionsRepository(db *sqlx.DB) *QuestionsRepository {
	return &QuestionsRepository{
		db: db,
	}
}

func (q *QuestionsRepository) CreateQuestion(type_ string, content string)  (int, error){
	var questionID int
	query := fmt.Sprintf("INSERT INTO %s (type, content) VALUES ($1, $2) RETURNING id", questionsTable)
	row := q.db.QueryRow(query, type_, content)
	if err := row.Scan(&questionID); err != nil {
		return 0, err
	}
	return questionID, nil
}
