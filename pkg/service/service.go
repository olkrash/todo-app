package service

import (
	todo_app "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (map[string]string, error)
}

type TodoList interface {
	Create(userId int, list todo_app.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
