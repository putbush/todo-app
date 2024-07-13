package service

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(authorization Authorization, todoList TodoList, todoItem TodoItem) *Service {
	return &Service{Authorization: authorization, TodoList: todoList, TodoItem: todoItem}
}
