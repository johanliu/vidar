package context

import "net/http"

type Parameter struct {
	Key   string
	Value string
}

type Parameters []Parameter

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Container      map[string]interface{}
	Parameters     Parameters
}

func (ctx *Context) Set(key string, value interface{}) error {
	if ctx.Container == nil {
		ctx.Container = make(map[string]interface{})
	}
	ctx.Container[key] = value

	return nil
}

func (ctx *Context) Get(key string) (interface{}, bool) {
	if ctx.Container != nil {
		value, exist := ctx.Container[key]
		return value, exist
	}

	return nil, false
}
