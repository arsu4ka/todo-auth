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
	var Todo models.Todo
	result := td.db.First(&Todo, id)
	return &Todo, result.Error
}

func (td *TodoService) FindByUser(userID uint) ([]*models.Todo, error) {
	var Todos []*models.Todo
	result := td.db.Where("user_id = ?", userID).Find(&Todos)
	return Todos, result.Error
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
