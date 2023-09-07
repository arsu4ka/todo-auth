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
	FindById(string) (*models.Todo, error)
	FindByUser(uint) ([]*models.Todo, error)
	Create(*models.Todo) error
	Update(string, *models.Todo) error
	Delete(string) error
}

type IVerificationService interface {
	Create(verif *models.Verification) error
	FindById(id uuid.UUID) (*models.Verification, error)
	FindByUserId(userID uint) (*models.Verification, error)
	Delete(id uuid.UUID) error
}

type IResetService interface {
	Create(reset *models.Reset) error
	FindById(id uuid.UUID) (*models.Reset, error)
	FindByUserId(userID uint) (*models.Reset, error)
	Delete(id uuid.UUID) error
}

type IEmailService interface {
	SendVerificationLink(toAdress, toName, token string) error
	SendResetLink(toAddress, toName, token string) error
}
