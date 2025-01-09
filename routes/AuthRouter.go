package routes

import (
	"ExpenseMate/controller"

	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	r.HandleFunc("/register", controller.RegisterController).Methods("POST")
	r.HandleFunc("/login", controller.LoginController).Methods("POST")
}
