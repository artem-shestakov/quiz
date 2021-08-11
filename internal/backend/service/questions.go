package service

import (
	"gitlab.com/artem-shestakov/quiz/internal/backend/repository"
	"gitlab.com/artem-shestakov/quiz/internal/model"
)

type QuestionsService struct {
	repository repository.Questions
}

func NewQuestionsService(repository repository.Questions) *QuestionsService {
	return &QuestionsService{
		repository: repository,
	}
}

func (q *QuestionsService) CreateQuestion(type_ string, content string)  (int, error) {
	return q.repository.CreateQuestion(type_, content)
}

func (q *QuestionsService) GetQuestion(questionId int) (model.Question, error) {
	return q.repository.GetQuestion(questionId)
}

func (q *QuestionsService) IsCorrectAnswerExist(questionId int)  bool {
	return q.repository.IsCorrectAnswerExist(questionId)
}