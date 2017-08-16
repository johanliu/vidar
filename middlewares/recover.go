package middlewares

import (
	"fmt"
	"net/http"
)

func RecoverHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		h.ServeHTTP(w, r)

	})
}
