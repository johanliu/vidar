package context

import (
	"encoding/json"
	"net/http"

	"github.com/johanliu/Vidar/logger"
)

type parameters []parameter

type parameter struct {
	key   string
	value string
}

type Context struct {
	Response
	parameters
	Request   *http.Request
	container map[string]interface{}
	status    int
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

func (ctx *Context) JSON(code int, body interface{}) {
	if err := ctx.SetStatus(code); err != nil {
		logger.Error.Panicf("Set status code failed: %v", err)
	}

	if err := ctx.SetContentType("application/json; charset = utf-8"); err != nil {
		logger.Error.Panicf("Set content type failed: %v", err)
	}

	if err := json.NewEncoder(ctx.ResponseWriter).Encode(body); err != nil {
		logger.Error.Panicf("Set payload body failed: %v", err)
	}
}

//TODO: find a generic function to handle render
/*
func (ctx *Context) Render(code int, body interface{}) {
	if err := ctx.SetStatus(code); err != nil {
		logger.Error.Panicln("Set return code failed")
	}

	t := reflect.TypeOf(body).String()
	switch t {
	case "[]uint8":
	case "string":
	default:
		logger.Info.Printf("This is %s ", t)
	}
}*/
