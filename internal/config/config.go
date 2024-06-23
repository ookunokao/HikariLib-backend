package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	App
	DB
}

type App struct {
	Env string `env:"APP_ENV" env-default:"prod"`
}

type DB struct {
	Service string `env:"DB_SERVICE" env-default:"postgres"`
	Host    string `env:"DB_HOST" env-default:"localhost"`
	Port    string `env:"DB_PORT" env-default:"5432"`
	Name    string `env:"DB_NAME" env-default:"postgres"`
	User    string `env:"DB_USER" env-default:"postgres"`
	Pass    string `env:"DB_PASS" env-default:"postgres"`
	SslMode string `env:"DB_SSL_MODE" env-default:"false"`
}

func MustLoad() Config {
	configPath := "./build/.env"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Config file not found")
	}

	var cnf Config
	if err := cleanenv.ReadConfig(configPath, &cnf); err != nil {
		log.Fatalf("cannot load config: %s", err)
	}

	return cnf
}
