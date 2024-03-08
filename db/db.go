package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Rhymond/go-money"
)

type Repeated int
type Tag string

const (
	None  Repeated = iota
	Daily
	Weekly
	BiWeekly
	Monthly
	Yearly
	Custom
)

const (
	Income  Tag = "income"
	Expense Tag = "expense"
	Balance Tag = "balance"
)

type BudgetEntry interface {
	RefreshDate()
}

type IncomeBudgetEntry struct {
	ID      	int
	Label   	string
	Amount 		*money.Amount
	Date  		time.Time
	Repeated  Repeated
	Tag 			Tag
}

type ExpenseBudgetEntry struct {
	ID      	int
	Label   	string
	Amount 		*money.Amount
	Date  		time.Time
	Repeated  Repeated
	Tag 			Tag
}

type BalanceBudgetEntry struct {
	ID 				int
	Label 		string
	Amount 		*money.Amount
	Date  		time.Time
	Tag 			Tag
}

func (i IncomeBudgetEntry) RefreshDate(data *time.Time) time.Time {
	var refreshedDate = time.Now().Year()
	res := time.Date(refreshedDate,data.Month(), data.Day(), 0, 0, 0, 0, time.Local)
	return res
}

func (i ExpenseBudgetEntry) RefreshDate(data *time.Time) time.Time {
	var refreshedDate = time.Now().Year()
	res := time.Date(refreshedDate,data.Month(), data.Day(), 0, 0, 0, 0, time.Local)
	return res
}

func (i BalanceBudgetEntry) RefreshDate(data *time.Time) time.Time {
	var refreshedDate = time.Now().Year()
	res := time.Date(refreshedDate,data.Month(), data.Day(), 0, 0, 0, 0, time.Local)
	return res
}

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

func RefreshDate(data *time.Time) time.Time {
	var refreshedDate = time.Now().Year()
	res := time.Date(refreshedDate,data.Month(), data.Day(), 0, 0, 0, 0, time.Local)
	return res
}

func Close(db *sql.DB) {
	db.Close()
	fmt.Println("Database connection closed")
}
