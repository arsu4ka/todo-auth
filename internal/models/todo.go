package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          string `gorm:"type:varchar(20);primaryKey"`
	Task        string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text"`
	Completed   bool   `gorm:"default:false;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	UserID uint
}

func (t *Todo) Validate() error {
	return nil
}

func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	fullUuid := uuid.New().String()
	splitted := strings.Split(fullUuid, "-")
	todoId := splitted[len(splitted)-1]
	t.ID = todoId
	return nil
}
