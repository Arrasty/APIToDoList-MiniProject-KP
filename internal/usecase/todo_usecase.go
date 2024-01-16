package usecase

import "github.com/Arrasty/api_todolist/internal/domain"

type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error) // Sesuaikan signature
}

type TodoUseCase interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error) // Sesuaikan signature
}

type todoUseCase struct {
	todoRepository TodoRepository
}

func NewTodoUseCase(todoRepository TodoRepository) TodoUseCase {
	return &todoUseCase{todoRepository}
}

func (uc *todoUseCase) Create(todo *domain.Todo) error {
	return uc.todoRepository.Create(todo)
}

func (uc *todoUseCase) GetAll() ([]domain.Todo, error) {
	return uc.todoRepository.GetAll()
}

func (uc *todoUseCase) GetByID(id uint) (*domain.Todo, error) {
	return uc.todoRepository.GetByID(id)
}

func (uc *todoUseCase) Update(todo *domain.Todo) error {
	return uc.todoRepository.Update(todo)
}

func (uc *todoUseCase) Delete(id uint) error {
	return uc.todoRepository.Delete(id)
}

func (uc *todoUseCase) MarkAsCompleted(id uint) error {
	return uc.todoRepository.MarkAsCompleted(id)
}

func (uc *todoUseCase) GetCompleted() ([]domain.Todo, error) {
	return uc.todoRepository.GetCompleted()
}