package services

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/google/uuid"
)

type IUserService interface {
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
}

type ITodoService interface {
	FindById(id string) (*models.Todo, error)
	FindByUser(userId uint, limit, page int) ([]*models.Todo, error)
	Create(todo *models.Todo) error
	Update(id string, updatedTodo *models.Todo) error
	Delete(id string) error
	GetTotalRecordCount() (int64, error)
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
