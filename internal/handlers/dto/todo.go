package dto

import "github.com/arsu4ka/todo-auth/internal/models"

type CreateTodoDto struct {
	Task      string `json:"task" binding:"required"`
	Completed bool   `json:"completed"`
}

type UpdateTodoDto struct {
	Task      string `json:"task" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
}

type ResponseTodoDto struct {
	ID        uint   `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func NewResponseTodoDto(todo *models.Todo) *ResponseTodoDto {
	return &ResponseTodoDto{
		ID:        todo.ID,
		Task:      todo.Task,
		Completed: todo.Completed,
	}
}
