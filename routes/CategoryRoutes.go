package routes

import (
	"ExpenseMate/controller"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	r.HandleFunc("/create-category", controller.CreateCategoryController).Methods("POST")
	r.HandleFunc("/get-category", controller.GetCategory).Methods("GET")
	r.HandleFunc("/get-category/{id}", controller.GetSingleCategoryController).Methods("GET")
	r.HandleFunc("/delete-category/{id}", controller.DeleteCategoryController).Methods("DELETE")
}
