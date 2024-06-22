package main

import (
	"github.com/spf13/viper"
	HikariLib_backend "github.com/yuminekosan/HikariLib-backend"
	"github.com/yuminekosan/HikariLib-backend/pkg/controller"
	"github.com/yuminekosan/HikariLib-backend/pkg/repository"
	"github.com/yuminekosan/HikariLib-backend/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}
	rep := repository.NewRepository()
	services := service.NewService(rep)
	routes := controller.NewRoutes(services)
	srv := new(HikariLib_backend.Server)
	if err := srv.Run(viper.GetString("port"), routes.InitRoutes()); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
