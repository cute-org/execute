package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lib/pq"
)

var DB *sql.DB

func IsUniqueViolation(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		return pgErr.Code == "23505"
	}
	return false
}

func InitDB() {
	var err error

	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST environment variable is required")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatal("DB_PORT environment variable is required")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER environment variable is required")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("DB_PASSWORD environment variable is required")
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		log.Fatal("DB_NAME environment variable is required")
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		log.Fatal("DB_SSLMODE environment variable is required")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        salt TEXT NOT NULL,
        passwordhash TEXT NOT NULL,
        avatar BYTEA
    );`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("failed to create table:", err)
	}

	// Configure the database connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)
}
