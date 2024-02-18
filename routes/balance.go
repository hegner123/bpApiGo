package routes

import (
	"bp/api/db"
	"net/http"
)

func balance( w http.ResponseWriter, r *http.Request){
		//write post Request to db
		//return balance
		//return error

		balance := db.GetBalance()
		if balance == nil {
			http.Error(w, "Error getting balance", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(balance)
		return


}