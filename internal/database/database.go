package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	DB_USER := os.Getenv("DB_USER")
	DB_PW := os.Getenv("DB_PW")
	DB_NAME := os.Getenv("DB_NAME")
	conn := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		DB_USER,
		DB_PW,
		DB_NAME,
	)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
