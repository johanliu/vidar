package router

import (
	"net/http"

	"github.com/johanliu/Vidar/logger"
)

type Router struct {
	handlers map[string][]*Endpoint
	NotFound http.Handler
}

type Endpoint struct {
	method string
	http.Handler
	redirect bool
}

func New() *Router {
	return &Router{handlers: make(map[string][]*Endpoint)}
}

func (r *Router) Add(method string, path string, h http.Handler) {

	if path[0] != '/' {
		logger.Error.Fatalf("Path must begin with '/' but in : %s", path)
	}

	r.handlers[path] = append(r.handlers[path], &Endpoint{method: method, Handler: h})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if eds, ok := r.handlers[req.URL.Path]; ok {
		for _, ed := range eds {
			if ed.method == req.Method {
				ed.Handler.ServeHTTP(w, req)
				return
			}
		}

		logger.Info.Printf("%s:%s Method Not Allowed", req.Method, req.URL.Path)
		http.Error(w, "Method Not Allowed", 405)

	} else {

		if r.NotFound != nil {
			r.NotFound.ServeHTTP(w, req)
			return
		}

		logger.Info.Printf("%s Not Found", req.URL.Path)
		http.Error(w, "URL Not Found", 404)
	}
}

func (r *Router) GET(path string, h http.Handler) {
	r.Add("GET", path, h)
}

func (r *Router) POST(path string, h http.Handler) {
	r.Add("POST", path, h)
}

func (r *Router) DELETE(path string, h http.Handler) {
	r.Add("DELETE", path, h)
}

func (r *Router) PUT(path string, h http.Handler) {
	r.Add("PUT", path, h)
}

// TODO: add regex or tree-based path resolution
