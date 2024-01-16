package repository

import (
	"time"
	"github.com/Arrasty/api_todolist/internal/domain"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error) // Perubahan pada baris ini
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error) // Tambahkan metode GetCompleted
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) GetAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) GetByID(id uint) (*domain.Todo, error) { // Perubahan pada baris ini
	var todo domain.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Todo{}, id).Error
}

func (r *todoRepository) MarkAsCompleted(id uint) error {
	return r.db.Model(&domain.Todo{}).Where("id = ?", id).Update("completed", true).Update("completed_at", time.Now()).Error
}

func (r *todoRepository) GetCompleted() ([]domain.Todo, error) {
	var completedTodos []domain.Todo
	err := r.db.Where("completed = ?", true).Find(&completedTodos).Error
	return completedTodos, err
}
