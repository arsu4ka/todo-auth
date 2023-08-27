package sqlservices

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResetService struct {
	db *gorm.DB
}

func NewResetService(db *gorm.DB) *ResetService {
	return &ResetService{
		db: db,
	}
}

func (rs *ResetService) Create(reset *models.Reset) error {
	return rs.db.Create(reset).Error
}

func (rs *ResetService) FindById(id uuid.UUID) (*models.Reset, error) {
	var reset models.Reset
	result := rs.db.First(&reset, id)
	return &reset, result.Error
}

func (rs *ResetService) FindByUserId(userID uint) (*models.Reset, error) {
	var reset models.Reset
	result := rs.db.Where("user_id = ?", userID).First(&reset)
	return &reset, result.Error
}

func (rs *ResetService) Delete(id uuid.UUID) error {
	return rs.db.Delete(&models.Verification{}, id).Error
}
