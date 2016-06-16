package context

func (ctx *Context) Status(code int) {
	ctx.ResponseWriter.WriteHeader(code)
}

func (ctx *Context) Header(key string, value string) {
	if len(value) == 0 {
		ctx.ResponseWriter.Header().Del(key)
	} else {
		ctx.ResponseWriter.Header().Set(key, value)
	}
}
