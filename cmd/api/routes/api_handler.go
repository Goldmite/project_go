package routes

import (
	"database/sql"

	"github.com/Goldmite/project_go/internal/auth"
	"github.com/Goldmite/project_go/internal/handlers"
	"github.com/Goldmite/project_go/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	authHandler := auth.NewAuthHandler()

	bookService := services.NewBookService(db)
	bookHandler := handlers.NewBookHandler(bookService)

	// Public routes
	public := router.Group("/api")
	{
		public.GET("/public", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Public endpoint - no auth required",
			})
		})
		public.GET("/login", authHandler.LoginRedirect)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(auth.EnsureValidToken())
	{
		protected.GET("/profile", authHandler.UserProfile)
		protected.GET("/protected-ping", authHandler.ProtectedPing)
		protected.POST("/books", bookHandler.CreateBookHandler)
		protected.GET("/books/:isbn", bookHandler.GetBookByIsbnHandler)
	}

	return router
}
