package middlewares

import (
	"net/http"

	"github.com/johanliu/Vidar/logger"
)

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO

		logger.Info.Printf("Enter [%s] %q", r.Method, r.URL.String())

		h.ServeHTTP(w, r)

		logger.Info.Printf("Leave [%s]", r.Method)
	})
}
