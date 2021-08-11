package handler

import (
	"encoding/json"
	"fmt"
	"gitlab.com/artem-shestakov/quiz/internal/model"
	"net/http"
)

func (h *Handler) CreateAnswers() http.HandlerFunc {
	var answer model.Answer
	return func(rw http.ResponseWriter, r *http.Request) {
		// Decode from JSON
		err := json.NewDecoder(r.Body).Decode(&answer)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if !h.canWeAddAnswer(answer) {
			fmt.Println("Correct answer for this question already exists")
			return
		}
		answer, err := h.service.CreateAnswer(answer.QuestionID, answer.Content, answer.IsCorrect)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		h.Response(rw, http.StatusOK, map[string]interface{}{
			"answer": answer,
		})

	}
}

func (h *Handler) canWeAddAnswer(answer model.Answer) bool {
	// Get questions info
	question, err := h.service.GetQuestion(answer.QuestionID)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	// If single_choice question already has one correct answer we can't add second
	if question.Type == "single_choice" && h.service.IsCorrectAnswerExist(answer.QuestionID) {
		return false
	}
	return true
}
