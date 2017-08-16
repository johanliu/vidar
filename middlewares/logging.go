package middlewares

import (
	"fmt"
	"net/http"
)

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("Enter [%s] %q", r.Method, r.URL.String())
		h.ServeHTTP(w, r)
		fmt.Printf("Leave [%s]", r.Method)
	})
}
