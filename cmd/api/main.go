package main

import (
	"github.com/Goldmite/project_go/cmd/api/routes"
	"github.com/Goldmite/project_go/internal/database"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	db, _ := database.Connect()

	defer db.Close()

	router := routes.SetupRoutes(db)

	router.Run(":3000")
}
