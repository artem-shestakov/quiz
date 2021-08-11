package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/artem-shestakov/quiz/internal/model"
)

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

func (q *QuestionsRepository) GetQuestion(questionId int) (model.Question, error)  {
	var question model.Question
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", questionsTable)
	err := q.db.Get(&question, query, questionId)
	return question, err
}

func (q *QuestionsRepository) IsCorrectAnswerExist(questionId int)  bool {
	var isCorrect bool
	query := fmt.Sprintf("SELECT is_correct FROM %s INNER JOIN %s a ON questions.id = a.question_id AND a.is_correct=true AND questions.id = $1;", questionsTable, answersTable)
	err := q.db.QueryRow(query).Scan(&isCorrect, questionId)
	if err != nil {
		return true
	}
	return false
}
