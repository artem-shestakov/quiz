package service

import "gitlab.com/artem-shestakov/quiz/internal/backend/repository"

type Questions interface {
	CreateQuestion(type_ string, content string)  (int, error)
}

type Service struct {
	Questions
}

func CreateService(repository *repository.Repository) *Service {
	return &Service{
		Questions: NewQuestionsService(repository.Questions),
	}
}

