package sqlservices

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (ur *UserService) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := ur.db.First(&user, id)
	return &user, result.Error
}

func (ur *UserService) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := ur.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (ur *UserService) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return ur.db.Create(user).Error
}

func (ur *UserService) Update(user *models.User) error {
	return ur.db.Save(user).Error
}
