package todo

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type TodoItems struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
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
