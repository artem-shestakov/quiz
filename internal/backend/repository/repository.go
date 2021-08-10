package repository

import "github.com/jmoiron/sqlx"

type Questions interface {
	CreateQuestion(type_ string, content string)  (int, error)
}

type Repository struct {
	Questions
}

func CreateRepository(db *sqlx.DB) * Repository {
	return &Repository{
		Questions: NewQuestionsRepository(db),
	}
}
