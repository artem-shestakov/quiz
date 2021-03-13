package server

import (
	"io"
	"net/http"
)

func (s *Server) Hello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Hello")
	}
}
