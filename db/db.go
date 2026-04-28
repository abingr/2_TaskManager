package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func ConnectDB() {
	var err error

	Conn, err = pgx.Connect(context.Background(),
		"postgres://postgres:postgres@localhost:5433/taskdb?sslmode=disable")

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	log.Println("Connected to PostgreSQL")
}
