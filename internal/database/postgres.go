package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Connect to the database
func Connect(connectionString string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return db, fmt.Errorf("could not open database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		return db, fmt.Errorf("could not ping database: %v", err)
	}

	log.Println("connected to database")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
