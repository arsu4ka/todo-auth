package services

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/google/uuid"
)

type IUserService interface {
	FindByID(uint) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	Create(*models.User) error
	Update(user *models.User) error
}

type ITodoService interface {
	FindById(uint) (*models.Todo, error)
	FindByUser(uint) ([]*models.Todo, error)
	Create(*models.Todo) error
	Update(uint, *models.Todo) error
	Delete(uint) error
}

type IVerificationService interface {
	Create(verif *models.Verification) error
	FindById(id uuid.UUID) (*models.Verification, error)
	FindByUserId(userID uint) (*models.Verification, error)
	Delete(id uuid.UUID) error
}

type IEmailService interface {
	SendVerificationLink(toAdress, toName, token string) error
}
