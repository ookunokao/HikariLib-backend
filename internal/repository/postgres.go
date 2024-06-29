package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/yuminekosan/hikariLibBackend/internal/config"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	userItemsTable  = "user_items"
	todoItemsTable  = "todo_items"
	listsItemsTable = "list_items"
)

func NewPostgresDb(cnf config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		cnf.Service, cnf.User, cnf.Pass, cnf.Host, cnf.Port, cnf.Name, cnf.SslMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
