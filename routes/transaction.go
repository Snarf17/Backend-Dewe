package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	TransRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransRepository)

	r.HandleFunc("/transaction", h.FindTransaction).Methods("GET")
	// r.HandleFunc("/users/{id}", h.GetUsers).Methods("GET")
	// r.HandleFunc("/users", h.CreateUsers).Methods("POST")
	// r.HandleFunc("/users/{id}", h.UpdateUsers).Methods("PATCH")
	// r.HandleFunc("/users/{id}", h.DeleteUsers).Methods("DELETE")

}
