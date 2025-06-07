package main

import (
	"github.com/Goldmite/project_go/internal/database"
	"github.com/Goldmite/project_go/internal/handlers"
	"github.com/Goldmite/project_go/internal/services"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	db, _ := database.Connect()

	defer db.Close()

	bookService := services.NewBookService(db)
	bookHandler := handlers.NewBookHandler(bookService)
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ready"})
		})
	}
	api.POST("/books", bookHandler.CreateBookHandler)
	api.GET("/books/:isbn", bookHandler.GetBookByIsbnHandler)
	api.POST("/users", userHandler.CreateUserHandler) // Register
	api.GET("/users/:id", userHandler.GetUserHandler)
	api.GET("/users", userHandler.GetUserHandler) // Login | by email and password

	router.Run(":3000")
}
