package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	authrouter := r.PathPrefix("/api/v1/auth").Subrouter()
	categoryroutes := r.PathPrefix("api/v1/category").Subrouter()
	expenseroutes := r.PathPrefix("api/v1/expense").Subrouter()

	routes.AuthRouter(authrouter)
	routes.CategoryRoutes(categoryroutes)
	routes.ExpenseRoutes(expenseroutes)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Server is live")
	if err != nil {
		log.Fatal(err)
	}
}
