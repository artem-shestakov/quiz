package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/artem-shestakov/quiz/internal/model"
)

type Questions interface {
	CreateQuestion(type_ string, content string)  (int, error)
	GetQuestion(questionId int) (model.Question, error)
	IsCorrectAnswerExist(questionId int)  bool
}

type Answers interface {
	CreateAnswer(questionID int, content string, IsCorrect bool)  (*model.Answer, error)
}

type Repository struct {
	Questions
	Answers
}

func CreateRepository(db *sqlx.DB) * Repository {
	return &Repository{
		Questions: NewQuestionsRepository(db),
		Answers: NewAnswerRepository(db),
	}
}
