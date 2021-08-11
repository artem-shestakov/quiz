package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/artem-shestakov/quiz/internal/backend/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func CreateHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoute()  *mux.Router{
	router := mux.NewRouter()
	router.Use(h.LoggingMiddleware)

	//HealthCheck
	router.Handle("/ping", h.Ping())

	apiVersionOne := router.PathPrefix("/api/v1").Subrouter()
	//Questions
	apiVersionOne.Handle("/question", h.CreateQuestion()).Methods("POST")

	//Answers
	apiVersionOne.Handle("/answer", h.CreateAnswers()).Methods("POST")
	return router
}

func (h *Handler) Response(writer http.ResponseWriter, code int, message interface{})  {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	msg, err := json.Marshal(message)
	if err != nil {
		return
	}
	writer.Write(msg)
}
