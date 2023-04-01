package store

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"gorm.io/gorm"
)

type ToDoRepo struct {
	db *gorm.DB
}

func (td *ToDoRepo) GetByUser(userID uint) ([]*models.ToDo, error) {
	var todos []*models.ToDo
	result := td.db.Where("user_id = ?", userID).Find(&todos)
	return todos, result.Error
}

func (td *ToDoRepo) Create(todo *models.ToDo) error {
	return td.db.Create(todo).Error
}
