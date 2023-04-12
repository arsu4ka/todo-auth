package resptypes

import "github.com/arsu4ka/todo-auth/internal/models"

type TodoResponse struct {
	Id        uint   `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func NewTodoResponse(todo *models.ToDo) *TodoResponse {
	return &TodoResponse{
		Id:        todo.ID,
		Task:      todo.Task,
		Completed: todo.Completed,
	}
}
