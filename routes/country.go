package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func CountryRoutes(r *mux.Router) {
	CountryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(CountryRepository)

	r.HandleFunc("/country", h.FindCountry).Methods("GET")
	r.HandleFunc("/country/{id}", h.GetCountry).Methods("GET")
	r.HandleFunc("/country", h.CreateCountry).Methods("POST")
	r.HandleFunc("/country/{id}", h.UpdateCountry).Methods("PATCH")
	// r.HandleFunc("/users/{id}", h.DeleteUsers).Methods("DELETE")

}
