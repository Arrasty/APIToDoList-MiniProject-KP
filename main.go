package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
	"os"

	"github.com/Arrasty/api_todolist/internal/config"
	"github.com/Arrasty/api_todolist/internal/delivery/http"
	"github.com/Arrasty/api_todolist/internal/domain"
	"github.com/Arrasty/api_todolist/internal/repository"
	"github.com/Arrasty/api_todolist/internal/usecase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	db.AutoMigrate(&domain.Todo{})

	todoRepository := repository.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUseCase(todoRepository)
	todoHandler := http.NewTodoHandler(todoUseCase)

	r.POST("/todos", todoHandler.Create)
	r.GET("/todos", todoHandler.GetAll)
	r.GET("/todos/:id", todoHandler.GetByID)
	r.PUT("/todos/:id", todoHandler.Update)
	r.DELETE("/todos/:id", todoHandler.Delete)
	r.PUT("/todos/:id/complete", todoHandler.MarkAsCompleted)
	r.GET("/completed", todoHandler.GetCompleted)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(r.Run(":" + port))
}