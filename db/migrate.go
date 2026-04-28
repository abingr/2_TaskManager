package db

import (
	"context"
	"log"
)

func RunMigrations() {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`

	taskTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE,
		user_id INT REFERENCES users(id)
	);`

	_, err := Conn.Exec(context.Background(), userTable)
	if err != nil {
		log.Fatal("Failed creating users table:", err)
	}

	_, err = Conn.Exec(context.Background(), taskTable)
	if err != nil {
		log.Fatal("Failed creating tasks table", err)
	}

	log.Println("Database migration complete")
}
