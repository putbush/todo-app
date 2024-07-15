package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) CreateList(list todo.TodoList, userId int) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", TodoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err = row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUserListQuery, userId, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TodoListPostgres) GetAllLists(userID int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", TodoListsTable, usersListsTable)
	if err := t.db.Select(&lists, query, userID); err != nil {
		return lists, err
	}

	return lists, nil
}

func (t *TodoListPostgres) GetListByID(listID, userID int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 and tl.id = $2", TodoListsTable, usersListsTable)
	if err := t.db.Get(&list, query, userID, listID); err != nil {
		return list, err
	}

	return list, nil

}
