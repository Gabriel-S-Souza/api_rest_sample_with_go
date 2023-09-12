package main

import (
	"net/http"

	"github.com/gabriel-s-souza/api_rest_sample_with_go/routes"
)

func main() {
	router := routes.LoadRoutes()
	http.ListenAndServe(":8000", router)
}
