package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func NewRDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn())
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %w", err)
	}

	return db, nil
}

func dsn() string {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbName)
}
