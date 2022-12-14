package routes

import (
	"BackEnd/handlers"
	"BackEnd/pkg/connection"
	"BackEnd/pkg/middleware"
	"BackEnd/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(connection.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/transaction-id", middleware.Auth(h.GetUserId)).Methods("GET")
	r.HandleFunc("/create-transaction", middleware.Auth(middleware.UploadProof(h.CreateTransaction))).Methods("POST")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.GetTransactionById)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
