package main

import (
	"log"
	"os"

	notemaker "github.com/abdrakhmanovzh/notemaker2.0"
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/handler"
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/repository"
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/service"
)

func main() {
	db, err := repository.Connect()
	if err != nil {
		log.Fatalf("problem with connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(notemaker.Server)
	notemaker.LoadEnvVariables()

	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
