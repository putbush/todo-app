package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type Item struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description" binding:"required"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int `json:"id"`
	ListID int `json:"listID"`
	ItemID int `json:"itemID"`
}

type UserList struct {
	Id     int `json:"id"`
	UserID int `json:"userID"`
	ListID int `json:"listID"`
}

type UpdateListInput struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
}

func (u UpdateListInput) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Done        *bool   `json:"done" db:"done"`
}

func (u UpdateItemInput) Validate() error {
	if u.Title == nil && u.Description == nil && u.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
