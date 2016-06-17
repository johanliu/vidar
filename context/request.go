package context

func (ctx *Context) Query(key string) string {
	value, _ := ctx.getQuery(key)
	return value
}

func (ctx *Context) QueryDefault(key string, defaultvalue string) string {
	value, ok := ctx.getQuery(key)
	if ok {
		return defaultvalue
	}
	return value
}

func (ctx *Context) getQuery(key string) (string, bool) {
	values, ok := ctx.Request.URL.Query()[key]
	if ok {
		if len(values) > 0 {
			return values[0], true
		}
	}

	return "", false
}

func (param *parameters) Params(key string) (string, bool) {
	return param.getParams(key)
}

func (params *parameters) getParams(key string) (string, bool) {
	for _, param := range *params {
		if param.key == key {
			return param.value, true
		}
	}

	return "", false
}
