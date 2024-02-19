package db

import (
	"database/sql"
	"log"
	"os"
)

// ConnectDB connects to the database
func ConnectDB() *sql.DB {
	db, err := sql.Open(os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	return db
}



