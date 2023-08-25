package sqlservices

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationService struct {
	db *gorm.DB
}

func NewVerificationService(db *gorm.DB) *VerificationService {
	return &VerificationService{
		db: db,
	}
}

func (vs *VerificationService) Create(verif *models.Verification) error {
	return vs.db.Create(verif).Error
}

func (vs *VerificationService) FindById(id uuid.UUID) (*models.Verification, error) {
	var verif models.Verification
	result := vs.db.First(&verif, id)
	return &verif, result.Error
}

func (vs *VerificationService) FindByUserId(userID uint) (*models.Verification, error) {
	var verif models.Verification
	result := vs.db.Where("user_id = ?", userID).First(&verif)
	return &verif, result.Error
}

func (vs *VerificationService) Delete(id uuid.UUID) error {
	return vs.db.Delete(&models.Verification{}, id).Error
}
