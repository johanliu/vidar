package router

import (
	"net/http"

	"github.com/johanliu/Vidar/logger"
)

type Router struct {
	handlers map[string][]*Endpoint
}

type Endpoint struct {
	method string
	http.Handler
	redirect bool
}

func New() *Router {
	return &Router{handlers: make(map[string][]*Endpoint)}
}

func (r *Router) Handler(method string, path string, h http.Handler) {
	if path[0] != '/' {
		logger.Error.Fatalf("Path must begin with '/' but in : %s", path)
	}

	r.Add(method, path, h)
}

func (r *Router) Add(method string, path string, h http.Handler) {
	r.handlers[path] = append(r.handlers[path], &Endpoint{method: method, Handler: h})
}

// TODO: add regex or tree-based path resolution
