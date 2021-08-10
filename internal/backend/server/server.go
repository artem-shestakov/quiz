package server

import "net/http"

type HttpServer struct {
	server *http.Server
}

func CreateServer(addr string, port string, handler http.Handler) *HttpServer {
	return &HttpServer{
		server: &http.Server{
			Addr: addr + ":" + port,
			Handler: handler,
		},
	}
}

func (s *HttpServer) Run() error{
	return s.server.ListenAndServe()
}
