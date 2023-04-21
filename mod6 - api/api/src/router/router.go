package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Build and return a router with routes setted
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.SetUp(r)
}
