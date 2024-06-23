package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
	hikarilibbackend "github.com/yuminekosan/hikariLibBackend"
	"github.com/yuminekosan/hikariLibBackend/internal/config"
	"github.com/yuminekosan/hikariLibBackend/internal/controller"
	"github.com/yuminekosan/hikariLibBackend/internal/repository"
	"github.com/yuminekosan/hikariLibBackend/internal/service"
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
	// todo: add slog-log and remake all logs after for slog
	db, err := repository.NewPostgresDb(cnf)
	if err != nil {
		log.Error("failed to connect to database: %s", err.Error())
	}
	rep := repository.NewRepository(db)
	services := service.NewService(rep)
	routes := controller.NewRoutes(services)
	srv := new(hikarilibbackend.Server)
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
