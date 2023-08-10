package models

type Todo struct {
	ID        uint
	Task      string
	Completed bool `gorm:"default:false"`
	UserID    uint
}
