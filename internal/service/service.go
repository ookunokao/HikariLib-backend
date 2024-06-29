package service

import (
	hikarilibbackend "github.com/yuminekosan/hikariLibBackend"
	"github.com/yuminekosan/hikariLibBackend/internal/repository"
)

type Authorization interface {
	CreateUser(User hikarilibbackend.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
