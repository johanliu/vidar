package plugins

import (
	"net/http"
	"time"

	"github.com/johanliu/vidar"
)

func LoggingHandler(h http.Handler) http.Handler {
	return vidar.ContextFunc(func(ctx *vidar.Context) {
		req := ctx.Request()
		resp := ctx.Response()
		ctx.Log.Info("Enter [%s] %q Host: %q", req.Method, req.URL.String(), req.Host)
		start := time.Now()
		defer func(time.Time) {
			ctx.Log.Info("Leave [%s] %q Elapse: %q %d",
				req.Method,
				req.URL.String(),
				time.Since(start),
				resp.Status())
		}(start)
		ctx.Call(h)
	})
}
