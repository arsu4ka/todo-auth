package models

import "github.com/google/uuid"

type Verification struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	UserID uint
}

func NewVerification(userID uint) *Verification {
	return &Verification{
		ID:     uuid.New(),
		UserID: userID,
	}
}
