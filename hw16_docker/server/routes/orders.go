package routes

import (
	"github.com/gorilla/mux"
	orderControllers "github.com/varik-08/golang-hw/hw16_docker/server/controllers/orders"
)

func SetupOrdersRoutes(router *mux.Router) {
	router.HandleFunc("/orders/by-user", orderControllers.GetOrdersByUser).Methods("GET")
	router.HandleFunc("/orders", orderControllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", orderControllers.DeleteOrder).Methods("DELETE")
}
