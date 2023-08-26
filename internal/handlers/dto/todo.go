package dto

import (
	"time"

	"github.com/arsu4ka/todo-auth/internal/models"
)

type CreateTodoDto struct {
	Task        string `json:"task" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type UpdateTodoDto struct {
	Task        string `json:"task" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed" binding:"required"`
}

type ResponseTodoDto struct {
	ID          uint   `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewResponseTodoDto(todo *models.Todo) *ResponseTodoDto {
	return &ResponseTodoDto{
		ID:          todo.ID,
		Task:        todo.Task,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
