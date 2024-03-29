package usecase

import (
	"github.com/Arrasty/api_todolist/internal/domain"
	"github.com/Arrasty/api_todolist/internal/repository"
)

//interface yang mendefinisikan kontrak untuk operasi CRUD pada entitas Todo.
//sebagai dependensi untuk use case
// type TodoRepository interface {
// 	Create(todo *domain.Todo) error
// 	GetAll() ([]domain.Todo, error)
// 	GetByID(id uint) (*domain.Todo, error)
// 	Update(todo *domain.Todo) error
// 	Delete(id uint) error
// 	MarkAsCompleted(id uint) error
// 	GetCompleted() ([]domain.Todo, error) // Sesuaikan signature
// 	GetUnCompleted() ([]domain.Todo, error)
// }

//interface yang mendefinisikan kontrak untuk operasi bisnis pada entitas Todo.
//sebagai abstraksi antarmuka bagi use case yang akan diimplementasikan
type TodoUseCase interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error)
	GetUnCompleted() ([]domain.Todo, error)
	SearchByTitle(title string) ([]*domain.Todo, error)
}

//implementasi konkret dari TodoUseCase.
//Struct menyimpan instance dari TodoRepository sebagai properti
type todoUseCase struct {
	todoRepository repository.TodoRepository
}

//Deklarasi fungsi NewTodoUseCase yang mengembalikan nilai bertipe TodoUseCase. 
//Fungsi ini menerima satu parameter, yaitu todoRepository bertipe TodoRepository
func NewTodoUseCase(todoRepository repository.TodoRepository) TodoUseCase {
	//Fungsi ini membuat instance baru dari todoUseCase dan menginisialisasi propertinya dengan todoRepository yang diteruskan sebagai parameter
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
	completedTodos, err := uc.todoRepository.GetCompleted()
    if err != nil {
        return nil, err
    }

    return completedTodos, nil
	//return uc.todoRepository.GetCompleted()
}

// GetUnCompleted retrieves all uncompleted todos.
func (uc *todoUseCase) GetUnCompleted() ([]domain.Todo, error) {
    // Implement the logic to retrieve uncompleted todos from your storage (e.g., database).
    // Return the list of uncompleted todos and any potential errors.

    // Example implementation:
    uncompletedTodos, err := uc.todoRepository.GetUnCompleted()
    if err != nil {
        return nil, err
    }

    return uncompletedTodos, nil
}

func (uc *todoUseCase) SearchByTitle(title string) ([]*domain.Todo, error) {
	return uc.todoRepository.SearchByTitle(title)
}
