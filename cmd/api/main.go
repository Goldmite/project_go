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
	groupService := services.NewGroupService(db)
	groupHandler := handlers.NewGroupHandler(groupService)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ready"})
		})
	}

	api.GET("/books/:isbn", bookHandler.FetchByIsbnFromApiHandler)
	api.POST("/users/signup", userHandler.CreateUserHandler) // Register
	api.GET("/users/:id", userHandler.GetUserHandler)
	api.POST("/users/login", userHandler.GetUserHandler) // Login | by email and password
	api.GET("/books/user/:id", bookHandler.GetAllUserBooksHandler)
	api.POST("/books", bookHandler.AddNewBookForUserHandler)
	api.POST("/groups", groupHandler.CreateGroupHandler)
	api.POST("/groups/invites", groupHandler.SendInvitesHandler)
	api.GET("/groups/invites/:id", userHandler.GetUserInvitesHandler)
	api.POST("/groups/invites/accept", groupHandler.AcceptInvitationHandler)
	api.DELETE("/groups/invites/decline", groupHandler.DeclineInvitationHandler)

	router.Run(":3000")
}
