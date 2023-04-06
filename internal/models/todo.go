package models

type ToDo struct {
	ID        uint
	Task      string
	Completed bool
	UserID    uint
}
