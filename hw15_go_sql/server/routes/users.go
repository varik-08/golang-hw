package routes

import (
	"github.com/gorilla/mux"
	userControllers "github.com/varik-08/golang-hw/hw15_go_sql/server/controllers/users"
)

func SetupUsersRoutes(router *mux.Router) {
	router.HandleFunc("/users", userControllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/orders-statistics", userControllers.GetUsersOrdersStatistics).Methods("GET")
	router.HandleFunc("/users", userControllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userControllers.UpdateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userControllers.DeleteUser).Methods("DELETE")
}
