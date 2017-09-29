package middlewares

import (
	"net/http"
)

func SlashHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path != "/" && path[len(path)-1] != '/' {
			path += "/"
			r.RequestURI = path
			r.URL.Path = path
		}
		h.ServeHTTP(w, r)
	})
}
