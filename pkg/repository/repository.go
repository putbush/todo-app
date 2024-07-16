package repository

import (
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

const (
	UsersTable      = "users"
	TodoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Authorization interface {
	CreateUser(user *todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
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

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
