package bpApiGo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/hegner123/bpApiGo/routes"
)

func main(){

    r := mux.NewRouter()
    r.HandleFunc("/api/v1/health", HealthHandler).Methods("GET")
    r.HandleFunc("/api/v1/balance", routes.GetBalance).Methods("GET")
    r.HandleFunc("/api/v1/income", routes.GetIncome).Methods("GET")
    r.HandleFunc("/api/v1/expense", routes.GetExpense).Methods("GET")

    http.ListenAndServe(":8080", r)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}
