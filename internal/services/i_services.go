package services

import "github.com/arsu4ka/todo-auth/internal/models"

type IUserService interface {
	FindByID(uint) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	Create(*models.User) error
}

type ITodoService interface {
	FindById(uint) (*models.Todo, error)
	FindByUser(uint) ([]*models.Todo, error)
	Create(*models.Todo) error
	Update(uint, *models.Todo) error
	Delete(uint) error
}
