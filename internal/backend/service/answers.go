package service

import (
	"gitlab.com/artem-shestakov/quiz/internal/backend/repository"
	"gitlab.com/artem-shestakov/quiz/internal/model"
)

type AnswersService struct {
	repository repository.Answers
}

func NewAnswersService(repository repository.Answers) *AnswersService {
	return &AnswersService{
		repository: repository,
	}
}

func (a *AnswersService) CreateAnswer(questionID int, content string, isCorrect bool)  (*model.Answer, error) {
	return a.repository.CreateAnswer(questionID, content, isCorrect)
}
