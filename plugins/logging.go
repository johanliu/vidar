package plugins

import (
	"net/http"

	"github.com/johanliu/vidar"
)

func LoggingHandler(h http.Handler) http.Handler {
	return vidar.ContextFunc(func(ctx *vidar.Context) {
		req := ctx.Request()

		ctx.Log.Info("Enter [%s] %q %q", req.Method, req.URL.String(), req.Host)
		ctx.Call(h)
		ctx.Log.Info("Leave [%s] %q", req.Method, req.URL.String())
	})
}
