package service

import (
	"gitlab.com/artem-shestakov/quiz/internal/backend/repository"
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

type Service struct {
	Questions
	Answers
}

func CreateService(repository *repository.Repository) *Service {
	return &Service{
		Questions: NewQuestionsService(repository.Questions),
		Answers: NewAnswersService(repository.Answers),
	}
}

