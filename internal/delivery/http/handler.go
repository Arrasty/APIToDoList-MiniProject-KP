package http

import (
	"strconv"

	"github.com/Arrasty/api_todolist/internal/domain"
	"github.com/Arrasty/api_todolist/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoHandler(todoUseCase usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{todoUseCase}
}

func (h *TodoHandler) Create(c *gin.Context) {
	// Menggunakan struct yang hanya berisi field selain Complete
	type TodoCreateInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	// Parsing body request ke struct TodoCreateInput
	var input TodoCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "error parsing request body"})
		return
	}

	// Membuat objek Todo dengan nilai default Completed: false
	todo := &domain.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}

	if err := h.todoUseCase.Create(todo); err != nil {
		c.JSON(500, gin.H{"error": "error creating todo"})
		return
	}

	c.JSON(201, gin.H{"message": "todo created successfully", "todo": todo})
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.todoUseCase.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "error getting todos"})
		return
	}

	c.JSON(200, gin.H{"todos": todos})
}

func (h *TodoHandler) GetCompleted(c *gin.Context) {
	completedTodos, err := h.todoUseCase.GetCompleted()
	if err != nil {
		c.JSON(500, gin.H{"error": "error getting completed todos"})
		return
	}

	if len(completedTodos) == 0 {
		c.JSON(200, gin.H{"message": "No completed todos found"})
		return
	}

	c.JSON(200, gin.H{"completedTodos": completedTodos})
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todo ID"})
		return
	}

	todo, err := h.todoUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(200, gin.H{"todo": todo})
}

func (h *TodoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todo ID"})
		return
	}

	// Dapatkan todo dari database berdasarkan ID
	existingTodo, err := h.todoUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	// Parse body request ke struct Todo
	var updateData domain.Todo
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "error parsing request body"})
		return
	}

	// Update atribut sesuai data yang diterima
	if updateData.Title != "" {
		existingTodo.Title = updateData.Title
	}
	if updateData.Description != "" {
		existingTodo.Description = updateData.Description
	}

	// Update todo
	if err := h.todoUseCase.Update(existingTodo); err != nil {
		c.JSON(500, gin.H{"error": "error updating todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo updated successfully", "todo": existingTodo})
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todo ID"})
		return
	}

	if err := h.todoUseCase.Delete(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "error deleting todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo deleted successfully"})
}

func (h *TodoHandler) MarkAsCompleted(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todo ID"})
		return
	}

	// Dapatkan todo dari database berdasarkan ID
	existingTodo, err := h.todoUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	// Menandai tugas sebagai selesai hanya dengan memperbarui atribut Completed
	existingTodo.Completed = !existingTodo.Completed

	// Update todo
	if err := h.todoUseCase.Update(existingTodo); err != nil {
		c.JSON(500, gin.H{"error": "error updating todo"})
		return
	}

	c.JSON(200, gin.H{"message": "todo marked as completed successfully", "todo": existingTodo})
}