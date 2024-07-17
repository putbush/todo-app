package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
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

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", UsersListsTable)
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

	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id "+
		"WHERE ul.user_id = $1", TodoListsTable, UsersListsTable)
	if err := t.db.Select(&lists, query, userID); err != nil {
		return lists, err
	}

	return lists, nil
}

func (t *TodoListPostgres) GetListByID(listID, userID int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 and tl.id = $2", TodoListsTable, UsersListsTable)
	if err := t.db.Get(&list, query, userID, listID); err != nil {
		return list, err
	}

	return list, nil
}

func (t *TodoListPostgres) DeleteListByID(listID, userID int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE ul.list_id=tl.id AND ul.list_id=$1 AND ul.user_id=$2", TodoListsTable, UsersListsTable)
	_, err := t.db.Exec(query, listID, userID)
	return err
}

func (t *TodoListPostgres) Update(listID, userID int, input todo.UpdateListInput) (todo.TodoList, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, input.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	args = append(args, listID, userID)

	var newList todo.TodoList
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d RETURNING tl.*",
		TodoListsTable, setQuery, UsersListsTable, argID, argID+1)
	if err := t.db.Get(&newList, query, args...); err != nil {
		return newList, err
	}
	return newList, nil
}
