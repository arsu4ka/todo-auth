package dto

import (
	"time"

	"github.com/arsu4ka/todo-auth/internal/models"
)

type CreateUserDto struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseUserDto struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"time"`
}

func NewResponseUserDto(user *models.User) *ResponseUserDto {
	return &ResponseUserDto{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
