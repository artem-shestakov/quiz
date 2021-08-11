package handler

import (
	"encoding/json"
	"fmt"
	"gitlab.com/artem-shestakov/quiz/internal/model"
	"net/http"
)

func (h *Handler) CreateQuestion() http.HandlerFunc {
	var question model.Question
	return func(rw http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&question)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		id, err := h.service.CreateQuestion(question.Type, question.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		h.Response(rw, http.StatusOK, map[string]interface{}{
			"question_id": id,
		})
	}
}
