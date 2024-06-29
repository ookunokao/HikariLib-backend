package repository

import (
	"github.com/jmoiron/sqlx"
	hikarilibbackend "github.com/yuminekosan/hikariLibBackend"
)

type Authorization interface {
	CreateUser(User hikarilibbackend.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
