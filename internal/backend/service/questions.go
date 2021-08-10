package service

import "gitlab.com/artem-shestakov/quiz/internal/backend/repository"

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
