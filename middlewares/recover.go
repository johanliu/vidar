package middlewares

import (
	"net/http"
	"runtime"
)

var stackSize = 4 << 10 //4KB

func RecoverHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, stackSize)
				runtime.Stack(stack, false)
				log.Info("\n %s %s \n", err, stack)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
