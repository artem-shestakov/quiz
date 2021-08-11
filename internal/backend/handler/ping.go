package handler

import (
	"net/http"
)

func (h *Handler) Ping() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		h.Response(writer, 200, map[string]interface{}{
			"ping": "pong",
		})
	}
}
