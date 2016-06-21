package context

import (
	"net/http"
	"net/url"
)

type Parameters []parameter

type parameter struct {
	key   string
	value string
}

type Context struct {
	Response
	Parameters
	Values    url.Values
	Request   *http.Request
	container map[string]interface{}
	status    int
}

func New(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response:   Response{ResponseWriter: w},
		Request:    r,
		Parameters: []parameter{},
	}
}

func (ctx *Context) Set(key string, value interface{}) error {
	if ctx.container == nil {
		ctx.container = make(map[string]interface{})
	}
	ctx.container[key] = value

	return nil
}

func (ctx *Context) Get(key string) (interface{}, bool) {
	if ctx.container != nil {
		value, exist := ctx.container[key]
		return value, exist
	}

	return nil, false
}
