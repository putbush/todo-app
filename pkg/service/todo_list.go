package service

import (
	todo "todo-app"
	"todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) CreateList(list todo.TodoList, userId int) (int, error) {
	return t.repo.CreateList(list, userId)
}

func (t *TodoListService) GetAllLists(userID int) ([]todo.TodoList, error) {
	return t.repo.GetAllLists(userID)
}

func (t *TodoListService) GetListByID(listID, userID int) (todo.TodoList, error) {
	return t.repo.GetListByID(listID, userID)
}

func (t *TodoListService) DeleteListByID(listID, userID int) error {
	return t.repo.DeleteListByID(listID, userID)
}

func (t *TodoListService) Update(listID, userID int, input todo.UpdateListInput) (todo.TodoList, error) {
	return t.repo.Update(listID, userID, input)
}
