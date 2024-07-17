package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	todo "todo-app"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) CreateItem(item todo.Item, listID int) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", TodoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err = row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	createListsItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", ListsItemsTable)
	_, err = tx.Exec(createListsItemQuery, listID, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TodoItemPostgres) GetAllItems(userID, listID int) ([]todo.Item, error) {
	var items []todo.Item

	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li on ti.id = li.item_id "+
		"INNER JOIN %s ul on ul.list_id=li.list_id WHERE li.list_id = $1 AND ul.user_id=$2", TodoItemsTable, ListsItemsTable, UsersListsTable)
	if err := t.db.Select(&items, query, listID, userID); err != nil {
		return items, err
	}

	return items, nil
}

func (t *TodoItemPostgres) GetItemByID(userID, itemID int) (todo.Item, error) {
	var item todo.Item

	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li on ti.id = li.item_id "+
		"INNER JOIN %s ul on ul.list_id=li.list_id WHERE ti.id=$1 AND ul.user_id=$2", TodoItemsTable, ListsItemsTable, UsersListsTable)
	if err := t.db.Get(&item, query, itemID, userID); err != nil {
		return item, err
	}

	return item, nil
}

func (t *TodoItemPostgres) Update(userID, itemID int, input todo.UpdateItemInput) (todo.Item, error) {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argID))
		args = append(args, input.Done)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	args = append(args, itemID, userID)

	var newItem todo.Item
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s ul, %s li "+
		"WHERE ti.id = li.item_id AND li.list_id=ul.list_id AND ti.id=$%d AND ul.user_id=$%d RETURNING ti.*",
		TodoItemsTable, setQuery, UsersListsTable, ListsItemsTable, argID, argID+1)
	if err := t.db.Get(&newItem, query, args...); err != nil {
		return newItem, err
	}
	return newItem, nil
}

func (t *TodoItemPostgres) DeleteItemByID(userID, itemID int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s ul, %s li "+
		"WHERE li.item_id=ti.id AND li.list_id=ul.list_id AND ti.id=$1 AND ul.user_id=$2", TodoItemsTable, UsersListsTable, ListsItemsTable)
	_, err := t.db.Exec(query, itemID, userID)
	return err
}
