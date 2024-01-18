package repository

import (
	"time"

	"github.com/Arrasty/api_todolist/internal/domain"
	"gorm.io/gorm"
)

//interface yang mendefinisikan kontrak untuk operasi CRUD pada entitas Todo
type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error) // Perubahan pada baris ini
	Update(todo *domain.Todo) error
	Delete(id uint) error
	MarkAsCompleted(id uint) error
	GetCompleted() ([]domain.Todo, error) // Tambahkan metode GetCompleted
	GetUnCompleted() ([]domain.Todo, error)
	SearchByTitle(title string) ([]*domain.Todo, error)
}

//Struct menyimpan instance dari database GORM sebagai properti.
//menyimpan koneksi database GORM
type todoRepository struct {
	db *gorm.DB
}

//Membuat instance baru dari todoRepository dengan koneksi database GORM, lalu return ke TodoRepository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

//Func create yang memiliki parameter todo yg mewakili entitas Todo
//todo yg akan disimpan ke database
//todo memiliki tipe data pointer domain.Todo, perubahan akan disimpan ke struct ini
func (r *todoRepository) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

//Func GetAll
func (r *todoRepository) GetAll() ([]domain.Todo, error) {
	// todos adalah slice dari struct Todo
	var todos []domain.Todo
	//&todos adalah alamat memori dari slice todos, sehingga data hasil query dapat langsung dimasukkan ke dalam slice tersebut
	err := r.db.Find(&todos).Error
	return todos, err
}

//Func GetById
func (r *todoRepository) GetByID(id uint) (*domain.Todo, error) {
	//var todo dengan tipe data domain.Todo
	var todo domain.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

//Func Update
func (r *todoRepository) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

//Func Delete
func (r *todoRepository) Delete(id uint) error {
	//hapus record pada tabel Todo dan dari id apa
	return r.db.Delete(&domain.Todo{}, id).Error
}

//Func MarkAsCompleted
func (r *todoRepository) MarkAsCompleted(id uint) error {
	//update pada model todo bedasarkan id, lalu update 2 atribut
	return r.db.Model(&domain.Todo{}).Where("id = ?", id).Updates(map[string]interface{}{
		"completed":    true,
		"completed_at": time.Now(),
	}).Error
}

//Func GetCompleted
func (r *todoRepository) GetCompleted() ([]domain.Todo, error) {
	var completedTodos []domain.Todo
	//ambil semua value yang bernilai true, lalu disimpan di completedTodos
	err := r.db.Where("completed = ?", true).Find(&completedTodos).Error
	return completedTodos, err
}

func (r *todoRepository) GetUnCompleted() ([]domain.Todo, error) {
    var unCompletedTodos []domain.Todo
    err := r.db.Where("completed = ?", false).Find(&unCompletedTodos).Error
    return unCompletedTodos, err
}

func (r *todoRepository) SearchByTitle(title string) ([]*domain.Todo, error) {
	var todos []*domain.Todo
	// Menggunakan Find dengan kondisi WHERE untuk melakukan pencarian berdasarkan judul.
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&todos).Error
	return todos, err
}
