package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	SetupUsersRoutes(router)
	SetupProductsRoutes(router)
	SetupOrdersRoutes(router)

	return router
}
