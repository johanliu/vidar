package vidar

import (
	"net/http"
	"net/url"
)

// type Parameters []parameter

type Parameters struct {
	key   string
	value map[int]string
}

type Context struct {
	Request *http.Request
	Response
	Parameters *Parameters
	Values     url.Values
	container  map[string]interface{}
	status     int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {

	return &Context{
		Request:    r,
		Response:   Response{ResponseWriter: w},
		Parameters: &Parameters{key: "pathParam", value: r.Context().Value("abc").(map[int]string)},
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
