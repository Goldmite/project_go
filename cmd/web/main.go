package main

import (
	"github.com/Goldmite/project_go/cmd/database"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	database.Connect()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ready"})
	})
	router.Run(":3000")
}
