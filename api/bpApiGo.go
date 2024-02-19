package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main(){

    router := mux.NewRouter()
    router.HandleFunc("/api/v1/health", HealthHandler).Methods("GET")
    router.HandleFunc("/api/v1/balance", GetBalance).Methods("GET")
    router.HandleFunc("/api/v1/income", GetIncome).Methods("GET")
    router.HandleFunc("/api/v1/expense", GetExpense).Methods("GET")



    http.ListenAndServe(":8080", router)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
    db := dbConnect()
    defer db.Close()
    // ...
}

func GetIncome(w http.ResponseWriter, r *http.Request) {
    db := dbConnect()
    defer db.Close()
    // ...
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
    db := dbConnect()
    defer db.Close()
    // ...
}



func dbConnect() *sql.DB {
    db, err := sql.Open(os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
    if err != nil {
        log.Fatalf("Error connecting to database: %s", err)
    }
    return db
}


