package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Represents all API routes
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

func SetUp(r *mux.Router) *mux.Router {
	for _, route := range routesUsers {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
