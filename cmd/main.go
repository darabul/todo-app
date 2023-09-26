package main

import (
	"log"

	"github.com/darabul/todo-app"
	"github.com/darabul/todo-app/pkg/handler"
	"github.com/darabul/todo-app/pkg/repository"
	"github.com/darabul/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http Server: %s", err.Error())
	}
}
