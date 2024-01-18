package main

import (
	"github.com/gin-gonic/gin" // menangani permintaan HTTP dan membangun API RESTful
	"github.com/joho/godotenv"

	"log"
	"os"

	"github.com/Arrasty/api_todolist/internal/config"
	"github.com/Arrasty/api_todolist/internal/delivery/http"
	"github.com/Arrasty/api_todolist/internal/domain"
	"github.com/Arrasty/api_todolist/internal/repository"
	"github.com/Arrasty/api_todolist/internal/usecase"
)

// dieksekusi sebelum fungsi main memuat variabel lingkungan
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//menginisialisasi router Gin
	r := gin.Default()

	//terhubung ke database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	//melakukan migrasi model Todo ke database
	db.AutoMigrate(&domain.Todo{})

	//menyiapkan rute HTTP
	todoRepository := repository.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUseCase(todoRepository)
	todoHandler := http.NewTodoHandler(todoUseCase)

	//mendefinisikan rute untuk operasi CRUD pada Todos
	apiV1 := r.Group("todos/api/v1/")
	{
	apiV1.POST("/create", todoHandler.Create)
	apiV1.GET("", todoHandler.GetAll)
	apiV1.GET("/:id", todoHandler.GetByID)
	apiV1.PUT("/update/:id", todoHandler.Update)
	apiV1.DELETE("/delete/:id", todoHandler.Delete)
	apiV1.PUT("/complete/:id", todoHandler.MarkAsCompleted)
	apiV1.GET("/completed", todoHandler.GetCompleted)
	apiV1.GET("/uncompleted", todoHandler.GetUnCompleted)
	apiV1.GET("/search/:title", todoHandler.SearchByTitle)
	}
	//menjalankan aplikasi Gin pada port 3000 sebagai default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(r.Run(":" + port))
}
