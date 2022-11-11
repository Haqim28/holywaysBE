package routes

import (
	"holyways/handlers"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transaction", h.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/{id}", h.GetTransactionByFund).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransactionByUser).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
}
