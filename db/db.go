package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		userId INTEGER,
		FOREIGN KEY(userId) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		eventId INTEGER,
		userId INTEGER,
		FOREIGN KEY(eventId) REFERENCES events(id),
		FOREIGN KEY(userId) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("could not create registrations table")
	}

}
