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

func (td *TodoService) FindById(id uint) (*models.Todo, error) {
	var Todo models.Todo
	result := td.db.First(&Todo, id)

	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	return nil, gorm.ErrRecordNotFound
	// }

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

func (td *TodoService) Update(id uint, updatedTodo *models.Todo) error {
	oldTodo, err := td.FindById(id)
	if err != nil {
		return err
	}

	updatedTodo.ID = oldTodo.ID
	return td.db.Save(updatedTodo).Error
}

func (td *TodoService) Delete(id uint) error {
	return td.db.Delete(&models.Todo{}, id).Error
}
