package models

import "time"

type Todo struct {
	ID          uint
	Task        string `gorm:"type:string;not null"`
	Description string `gorm:"type:text"`
	Completed   bool   `gorm:"default:false;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	UserID uint
}

func (t *Todo) Validate() error {
	return nil
}
