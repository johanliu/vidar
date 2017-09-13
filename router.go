package vidar

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/johanliu/Vidar/constant"
)

//TODO: Should implement trie-based router data structure

type Node struct {
}

type Tree struct {
}

type Router struct {
	handlers  map[string][]*Endpoint
	ehandlers map[string][]*Endpoint
	NotFound  http.Handler
}

type Endpoint struct {
	method string
	http.Handler
	redirect  bool
	PathParam map[int]string
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string][]*Endpoint)}
}

func (r *Router) Add(method string, path string, h http.Handler) {
	if path[0] != '/' {
		fmt.Printf("Path must begin with '/' but in : %s", path)
	}

	ed := &Endpoint{
		method:    method,
		Handler:   h,
		PathParam: r.pathParamSplit(path),
	}

	r.handlers[path] = append(r.handlers[path], ed)
}

func (r *Router) pathParamSplit(path string) map[int]string {
	container := make(map[int]string)

	segment := strings.Split(path, "/")
	for index, item := range segment {
		if strings.HasPrefix(item, ":") {
			container[index] = strings.Split(item, ":")[1]
		}
	}

	// fmt.Println(container)

	return container
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if eds, ok := r.handlers[req.URL.Path]; ok {
		for _, ed := range eds {
			if ed.method == req.Method {
				abc := make(map[int]string)
				abc[2] = "a"

				ctx := context.WithValue(req.Context(), "abc", abc)
				// fmt.Println(ctx)

				ed.Handler.ServeHTTP(w, req.WithContext(ctx))
				return
			}
		}

		log.Error(constant.MethodNotAllowedError)

	} else {
		if r.NotFound != nil {
			r.NotFound.ServeHTTP(w, req)
			return
		}
		fmt.Printf("%s Not Found", req.URL.Path)
		http.Error(w, "URL Not Found", 404)
	}
}

func (r *Router) GET(path string, h http.Handler) {
	r.Add(constant.GET, path, h)
}

func (r *Router) POST(path string, h http.Handler) {
	r.Add(constant.POST, path, h)
}

func (r *Router) DELETE(path string, h http.Handler) {
	r.Add(constant.DELETE, path, h)
}

func (r *Router) PUT(path string, h http.Handler) {
	r.Add(constant.PUT, path, h)
}

func (r *Router) PATCH(path string, h http.Handler) {
	r.Add(constant.PATCH, path, h)
}
