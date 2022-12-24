package main

import (
	database "dewetour/databases"
	"dewetour/pkg/mysql"
	"dewetour/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Initial DB
	mysql.DatabaseInit()

	// Run Migration
	database.RunMigration()

	// r := mux.NewRouter()
	r := mux.NewRouter()

	// routes.Roue
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("Server Running in localhost:9000")
	http.ListenAndServe("localhost:9000", r)
}
