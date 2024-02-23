package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to db.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		price FLOAT NOT NULL,
		status TEXT NOT NULL,
		createdAt DATETIME NOT NULL,
		store_id INTEGER NOT NULL 
	)
	`
	_, err := DB.Exec(createProductsTable)
	if err != nil {
		panic("Could not initialize db tables.")
	}

}