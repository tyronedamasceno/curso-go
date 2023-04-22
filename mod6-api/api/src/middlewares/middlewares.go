package middlewares

import (
	"fmt"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authenticating...")
		next(w, r)
	}
}
