package handler

import (
	"log"
	"net/http"
)

type (
	//responseData for ResponseWriter implementation
	responseData struct {
		status int
		size   int
	}

	//loggingResponseWriter is ResponseWriter implementation
	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

//Write response
func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b) // write response using original http.ResponseWriter
	r.responseData.size += size // save size
	return size, err
}
// WriteHeader sends an HTTP response header with the provided status code.
func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode) // write status code using original http.ResponseWriter
	r.responseData.status = statusCode // save status code
}

// LoggingMiddleware logging http requests
func (h *Handler) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lrw := loggingResponseWriter{
			ResponseWriter: rw, // compose original http.ResponseWriter
			responseData: responseData,
		}
		next.ServeHTTP(&lrw, r)

		log.Printf("-- %s %s %s %d %d", r.RemoteAddr, r.Method, r.RequestURI, responseData.status, responseData.size)
	})
}
