package store

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (ur *UserRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := ur.db.First(&user, id)
	return &user, result.Error
}

func (ur *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := ur.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (ur *UserRepo) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return ur.db.Create(user).Error
}
