package routes

import (
	"github.com/gorilla/mux"
	productControllers "github.com/varik-08/golang-hw/hw15_go_sql/server/controllers/products"
)

func SetupProductsRoutes(router *mux.Router) {
	router.HandleFunc("/products", productControllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", productControllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productControllers.UpdateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productControllers.DeleteProduct).Methods("DELETE")
}
