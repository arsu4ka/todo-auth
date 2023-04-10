package models

type ToDo struct {
	ID        uint
	Task      string `json:"task" binding:"required"`
	Completed bool   `json:"completed"`
	UserID    uint
}
