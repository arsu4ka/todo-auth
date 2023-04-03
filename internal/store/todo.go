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

func (td *ToDoRepo) UpdateStatus(id uint, setTo bool) error {
	result := td.db.Model(&models.ToDo{}).Where("id = ?", id).Update("completed", setTo)
	return result.Error
}

func (td *ToDoRepo) UpdateFull(todo *models.ToDo, id uint) error {
	var oldTodo models.ToDo
	result := td.db.First(&oldTodo, id)
	if result.Error != nil {
		return result.Error
	}

	oldTodo.Task = todo.Task
	oldTodo.Completed = todo.Completed
	return td.db.Save(&oldTodo).Error
}
