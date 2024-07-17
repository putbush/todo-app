package repository

import (
	"github.com/jmoiron/sqlx"
	todo "todo-app"
	"todo-app/pkg/repository/postgres"
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
	CreateItem(item todo.Item, listID int) (int, error)
	GetAllItems(userID, listID int) ([]todo.Item, error)
	GetItemByID(userID, itemID int) (todo.Item, error)
	Update(userID, itemID int, input todo.UpdateItemInput) (todo.Item, error)
	DeleteItemByID(userID, itemID int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		TodoList:      postgres.NewTodoListPostgres(db),
		TodoItem:      postgres.NewTodoItemPostgres(db),
	}
}
