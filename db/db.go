package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	createTable()
	fmt.Println("Successfully created task table")
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
