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

// IsUniqueViolation reports whether an error is a PostgreSQL unique violation.
func IsUniqueViolation(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		return pgErr.Code == "23505"
	}
	return false
}

// InitDB initializes the global DB handle, creates required tables, and sets up retry logic on connection.
func InitDB() {
	// Read configuration from environment
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

	// Build connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Open database handle
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	// Retry ping until successful
	for {
		err = DB.Ping()
		if err == nil {
			break
		}
		log.Printf("failed to connect to database: %v; retrying in 10 seconds...", err)
		time.Sleep(10 * time.Second)
	}
	log.Println("successfully connected to database")

	// Create tables if they do not exist
	createUsers := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        salt TEXT NOT NULL,
        passwordhash TEXT NOT NULL,
        display_name VARCHAR(255),
        phone VARCHAR(20),
        birth_date DATE,
        role VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	if _, err := DB.Exec(createUsers); err != nil {
		log.Fatal("failed to create users table:", err)
	}

	createGroups := `
    CREATE TABLE IF NOT EXISTS groups (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        code TEXT NOT NULL UNIQUE,
        creator_user_id INTEGER NOT NULL REFERENCES users(id)
    );`
	if _, err := DB.Exec(createGroups); err != nil {
		log.Fatal("failed to create groups table:", err)
	}

	alterUsersGroup := `
    ALTER TABLE users
    ADD COLUMN IF NOT EXISTS group_id INTEGER REFERENCES groups(id) ON DELETE SET NULL;`
	if _, err := DB.Exec(alterUsersGroup); err != nil {
		log.Fatal("failed to alter users table to add group_id:", err)
	}

	alterUsersAvatar := `
    ALTER TABLE users
    ADD COLUMN IF NOT EXISTS avatar BYTEA;`
	if _, err := DB.Exec(alterUsersAvatar); err != nil {
		log.Fatal("failed to alter users table to add avatar column:", err)
	}

	createTasks := `
    CREATE TABLE IF NOT EXISTS tasks (
        id               SERIAL      PRIMARY KEY,
        group_id         INTEGER     NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        creator_user_id  INTEGER     NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
        creation_date    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        due_date         TIMESTAMPTZ NOT NULL,
        name             TEXT        NOT NULL,
        description      TEXT,
        points_value     INTEGER     NOT NULL,
        step             INTEGER     NOT NULL DEFAULT 1
    );`
	if _, err := DB.Exec(createTasks); err != nil {
		log.Fatal("failed to create tasks table:", err)
	}

	// Configure the database connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)
}
