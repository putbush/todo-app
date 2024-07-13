package repository

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(authorization Authorization, todoList TodoList, todoItem TodoItem) *Repository {
	return &Repository{Authorization: authorization, TodoList: todoList, TodoItem: todoItem}
}
