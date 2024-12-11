// repository/todo_repository.go
package repository

import (
	"gorm.io/gorm"
	"todo-app/models"
)

type TodoRepository interface {
	Create(todo *models.Todo) (*models.Todo, error)
	FindByID(id uint) (*models.Todo, error)
	FindAll() ([]models.Todo, error)
	Update(todo *models.Todo) (*models.Todo, error)
	Delete(id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

// Create実装
func (r *todoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

// FindByID実装
func (r *todoRepository) FindByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// FindAll実装
func (r *todoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// Update実装
func (r *todoRepository) Update(todo *models.Todo) (*models.Todo, error) {
	if err := r.db.Save(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

// Delete実装
func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Todo{}, id).Error
}
