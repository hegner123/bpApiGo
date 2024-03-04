package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/hegner123/bpApiGo/db"
)


func main(){
    fmt.Println("Starting the application...")
    
    db.DbInit()


    router := mux.NewRouter()
    router.HandleFunc("/api/v1/health", HealthHandler).Methods("GET")
    router.HandleFunc("/api/v1/balance", GetBalance).Methods("GET")
    router.HandleFunc("/api/v1/income", GetIncome).Methods("GET")
    router.HandleFunc("/api/v1/expense", GetExpense).Methods("GET")

    err := http.ListenAndServe(":8080", router)
    if err != nil {
        fmt.Println(err)
    }
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, err := w.Write([]byte("OK"))
    if err != nil {
        fmt.Println(err)
    }

    
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
    sqlDataBase := db.DbConnect()
    defer db.Close(sqlDataBase)
    // ...
}

func GetIncome(w http.ResponseWriter, r *http.Request) {
    sqlDataBase := db.DbConnect()
    defer db.Close(sqlDataBase)
    // ...
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
    sqlDataBase := db.DbConnect()
    defer db.Close(sqlDataBase)
    // ...
}






