package models

import "github.com/google/uuid"

type Reset struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uint
}

func NewReset(userID uint) *Reset {
	return &Reset{
		ID:     uuid.New(),
		UserID: userID,
	}
}
