package routes

import (
	"api/src/middlewares"
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
	routes := routesUsers
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequiresAuth {
			r.HandleFunc(route.URI, middlewares.Authenticate(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		}
	}

	return r
}
