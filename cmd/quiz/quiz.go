package main

import (
	"gitlab.com/artem-shestakov/quiz/internal/backend/handler"
	"gitlab.com/artem-shestakov/quiz/internal/backend/repository"
	"gitlab.com/artem-shestakov/quiz/internal/backend/server"
	"gitlab.com/artem-shestakov/quiz/internal/backend/service"
	"log"
)

func main() {
	db, err := repository.CreateDB(&repository.Database{
		Address: "127.0.0.1",
		Port: "5432",
		User: "quiz",
		Password: "quiz",
		DBName: "quiz",
		SSLMode: "disable",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	httpRepository := repository.CreateRepository(db)
	httpService := service.CreateService(httpRepository)
	httpHandler := handler.CreateHandler(httpService)
	srv := server.CreateServer("0.0.0.0", "5000", httpHandler.InitRoute())
	if err:= srv.Run(); err != nil {
		panic(err.Error())
	}
}
