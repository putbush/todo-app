package service

import (
	todo "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user *todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(list todo.TodoList, userId int) (int, error)
	GetAllLists(userID int) ([]todo.TodoList, error)
	GetListByID(listID, userID int) (todo.TodoList, error)
	DeleteListByID(listID, userID int) error
	Update(listID, userID int, input todo.UpdateListInput) (todo.TodoList, error)
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
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
