package main

import (
	"github.com/Goldmite/project_shelf/internal/database"
	"github.com/Goldmite/project_shelf/internal/handlers"
	"github.com/Goldmite/project_shelf/internal/services"
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
	statsService := services.NewStatsService(db)
	statsHandler := handlers.NewStatsHandler(statsService)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ready"})
		})
	}

	api.GET("/users/:email", userHandler.GetUserHandler)
	api.POST("/users/signup", userHandler.CreateUserHandler) // Register
	api.POST("/users/login", userHandler.GetUserHandler)     // Login | by email and password
	api.GET("/users/groups/:id", userHandler.GetGroupMembersHandler)
	api.GET("/users/invites/:id", userHandler.GetUserInvitesHandler)

	api.POST("/books", bookHandler.AddNewBookForUserHandler)
	api.GET("/books/:isbn", bookHandler.GetBookByIsbnHandler)
	api.GET("/books/user/:id", bookHandler.GetAllUserBooksHandler)
	api.GET("/books/groups/:id", bookHandler.GetAllGroupBooksHandler)

	api.POST("/groups", groupHandler.CreateGroupHandler)
	api.GET("/groups/:id", groupHandler.GetAllUserGroupsHandler)
	api.POST("/groups/invites", groupHandler.SendInvitesHandler)
	api.GET("/groups/invites/:id", groupHandler.GetInvitesHandler)
	api.POST("/groups/invites/accept", groupHandler.AcceptInvitationHandler)
	api.DELETE("/groups/invites/decline", groupHandler.DeclineInvitationHandler)

	api.PUT("/stats/progress/book", statsHandler.UpdateBookProgressHandler)
	api.GET("/stats/progress/book", statsHandler.GetBookProgressHandler)
	api.GET("/stats/user/:id", statsHandler.GetUserStatsHandler)

	router.Run(":3000")
}
