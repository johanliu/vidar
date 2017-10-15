package plugins

import (
	"net/http"

	"github.com/johanliu/vidar"
)

func AuthHandler(h http.Handler) http.Handler {
	return vidar.ContextFunc(func(ctx *vidar.Context) {
		// TODO: authentication
		ctx.Call(h)
	})
}
