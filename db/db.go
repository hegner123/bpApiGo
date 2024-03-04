package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DbInit() {
	fmt.Println("Connecting to database...")
	DbConnect()
	
}

func DbConnect() *sql.DB {
	db, err := sql.Open(os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	if err != nil {
		fmt.Println("Error connecting to database")
		log.Fatalf("Error connecting to database: %s", err)
	}
	fmt.Println("Connected to database")
	return db
}

func Close(db *sql.DB) {
	db.Close()
	fmt.Println("Database connection closed")
}



