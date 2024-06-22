package main

import (
	HikariLib_backend "github.com/yuminekosan/HikariLib-backend"
	"github.com/yuminekosan/HikariLib-backend/pkg/controller"
	"github.com/yuminekosan/HikariLib-backend/pkg/repository"
	"github.com/yuminekosan/HikariLib-backend/pkg/service"
	"log"
)

func main() {
	rep := repository.NewRepository()
	services := service.NewService(rep)
	routes := controller.NewRoutes(services)
	srv := new(HikariLib_backend.Server)
	if err := srv.Run("8001", routes.InitRoutes()); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
