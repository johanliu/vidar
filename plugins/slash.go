package plugins

import (
	"net/http"

	"github.com/johanliu/vidar"
)

func SlashHandler(h http.Handler) http.Handler {
	return vidar.ContextFunc(func(ctx *vidar.Context) {
		req := ctx.Request()

		path := req.URL.Path
		if path != "/" && path[len(path)-1] != '/' {
			path += "/"
			req.RequestURI = path
			req.URL.Path = path
		}

		ctx.Call(h)
	})
}
