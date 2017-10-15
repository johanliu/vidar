package plugins

import (
	"net/http"
	"runtime"

	"github.com/johanliu/vidar"
)

var stackSize = 4 << 10 //4KB

func RecoverHandler(h http.Handler) http.Handler {
	return vidar.ContextFunc(func(ctx *vidar.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, stackSize)
				runtime.Stack(stack, false)
				ctx.Log.Info("\n %s %s \n", err, stack)
				ctx.Error(vidar.InternalServerError)
			}
		}()
		ctx.Call(h)
	})
}
