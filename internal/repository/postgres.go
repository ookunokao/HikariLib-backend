package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yuminekosan/HikariLib-backend/internal/config"
)

func NewPostgresDb(cnf config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.User, cnf.Name, cnf.Pass, cnf.SslMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
