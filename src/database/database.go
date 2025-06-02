package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() {
	db, err := sql.Open("sqlite", "./database/bookshelf.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	fmt.Println("Connected to the SQLite database successfully.")

	sqlBytes, _ := os.ReadFile("database/schema.sql")
	db.Exec(string(sqlBytes))
}
