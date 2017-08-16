package middlewares

import (
	"fmt"
	"net/http"
)

func AuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("Hello")
		h.ServeHTTP(w, r)
		fmt.Printf("Bye")
	})
}
