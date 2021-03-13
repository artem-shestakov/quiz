package main

import (
	"log"

	"github.com/artem-shestakov/quiz/internal/quiz/server"
)

func main() {
	config := server.NewConfig()
	server := server.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
