package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUsers).Methods("GET")
	r.HandleFunc("/users", h.CreateUsers).Methods("POST")
	r.HandleFunc("/users/{id}", h.UpdateUsers).Methods("PATCH")
	r.HandleFunc("/users/{id}", h.DeleteUsers).Methods("DELETE")

}
