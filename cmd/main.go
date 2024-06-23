package main

import (
	"github.com/spf13/viper"
	HikariLib_backend "github.com/yuminekosan/HikariLib-backend"
	"github.com/yuminekosan/HikariLib-backend/internal/config"
	"github.com/yuminekosan/HikariLib-backend/internal/controller"
	"github.com/yuminekosan/HikariLib-backend/internal/repository"
	"github.com/yuminekosan/HikariLib-backend/internal/service"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	cnf := config.MustLoad()
	log := setupLogger(cnf.Env)
	log.Info("pizda")
	// todo: add slog-log and remake all logs after for slog
	db, err := repository.NewPostgresDb(cnf)
	if err != nil {
		log.Error("failed to connect to database: %s", err.Error())
	}
	rep := repository.NewRepository(db)
	services := service.NewService(rep)
	routes := controller.NewRoutes(services)
	srv := new(HikariLib_backend.Server)
	if err := srv.Run(viper.GetString("port"), routes.InitRoutes()); err != nil {
		log.Error("Error starting server: %s", err.Error())
	}
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return logger
}
