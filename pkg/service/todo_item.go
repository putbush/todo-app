package service

import (
	todo "todo-app"
	"todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (t *TodoItemService) CreateItem(item todo.Item, listID, userID int) (int, error) {
	_, err := t.listRepo.GetListByID(listID, userID)

	if err != nil {
		return 0, err
	}

	return t.repo.CreateItem(item, listID)
}

func (t *TodoItemService) GetAllItems(userID, listID int) ([]todo.Item, error) {
	return t.repo.GetAllItems(userID, listID)
}

func (t *TodoItemService) GetItemByID(userID, itemID int) (todo.Item, error) {
	return t.repo.GetItemByID(userID, itemID)
}

func (t *TodoItemService) Update(userID, itemID int, input todo.UpdateItemInput) (todo.Item, error) {
	return t.repo.Update(userID, itemID, input)
}

func (t *TodoItemService) DeleteItemByID(itemID, userID int) error {
	return t.repo.DeleteItemByID(itemID, userID)
}
