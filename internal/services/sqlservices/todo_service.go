package sqlservices

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{
		db: db,
	}
}

func (td *TodoService) FindById(id string) (*models.Todo, error) {
	todo := models.Todo{ID: id}
	result := td.db.First(&todo)
	return &todo, result.Error
}

func (td *TodoService) FindByUser(userID uint, limit, page int) ([]*models.Todo, error) {
	var offset int
	if page <= 0 {
		offset = -1
	} else {
		offset = (page - 1) * limit
	}
	if limit == 0 {
		limit = -1
	}

	var todos []*models.Todo
	result := td.db.Where("user_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Order("updated_at DESC").
		Find(&todos)
	return todos, result.Error
}

func (td *TodoService) Create(todo *models.Todo) error {
	if err := todo.Validate(); err != nil {
		return err
	}

	return td.db.Create(todo).Error
}

func (td *TodoService) Update(id string, updatedTodo *models.Todo) error {
	oldTodo, err := td.FindById(id)
	if err != nil {
		return err
	}

	oldTodo.Task = updatedTodo.Task
	oldTodo.Description = updatedTodo.Description
	oldTodo.Completed = updatedTodo.Completed
	return td.db.Save(oldTodo).Error
}

func (td *TodoService) Delete(id string) error {
	return td.db.Delete(&models.Todo{}, id).Error
}

func (td *TodoService) GetTotalRecordCount() (int64, error) {
	var count int64
	result := td.db.Model(&models.Todo{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
