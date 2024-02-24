package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db?_foreign_keys=on")
	
	if err != nil {
		panic("Could not connect to db.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		status TEXT NOT NULL,
		createdAt DATETIME NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not initialize User table.")
	}

	createStoresTable := `
	CREATE TABLE IF NOT EXISTS stores (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		status TEXT NOT NULL,
		createdAt DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)
	`
	_, err = DB.Exec(createStoresTable)
	if err != nil {
		panic("Could not initialize Store table.")
	}

	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		price FLOAT NOT NULL,
		status TEXT NOT NULL,
		createdAt DATETIME NOT NULL,
		store_id INTEGER NOT NULL,
		FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE
	)
	`
	_, err = DB.Exec(createProductsTable)
	if err != nil {
		panic("Could not initialize Product table.")
	}


}