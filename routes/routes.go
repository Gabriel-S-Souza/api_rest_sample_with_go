package routes

import (
	"github.com/gabriel-s-souza/api_rest_sample_with_go/controllers"
	"github.com/gorilla/mux"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.GetHello).Methods("GET")
	router.HandleFunc("/product", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProductById).Methods("GET")
	return router
}
