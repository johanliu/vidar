package middlewares

import (
	"net/http"

	"github.com/johanliu/mlog"
)

var log = mlog.NewLogger()

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Info("Enter [%s] %q %q", r.Method, r.URL.String(), r.Host)
		h.ServeHTTP(w, r)
		log.Info("Leave [%s] %q", r.Method, r.URL.String())
	})
}
