package vidar

import (
	"net/http"
	"net/url"
	"strings"
)

type Router struct {
	tree     *node
	NotFound http.Handler
}

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]http.Handler
}

func (n *node) addNode(method, path string, h http.Handler) {
	components := strings.Split(path, "/")[1:]
	count := len(components)

	for {
		if count == 0 {
			break
		}

		oldNode, component := n.findNode(components, nil)
		if oldNode.component == component && count == 1 {
			oldNode.methods[method] = h
			return
		}

		newNode := &node{
			component: component,
			methods:   make(map[string]http.Handler),
		}

		if len(component) > 0 && component[0] == ':' {
			newNode.isNamedParam = true
		} else {
			newNode.isNamedParam = false
		}

		if count == 1 {
			newNode.methods[method] = h
		}

		oldNode.children = append(oldNode.children, newNode)
		count--
	}
}

func (n *node) findNode(components []string, params url.Values) (*node, string) {
	component := components[0]
	if len(n.children) > 0 {

		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.component[1:], component)
				}
				next := components[1:]
				if len(next) > 0 {
					return child.findNode(next, params)
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}

/*
type Endpoint struct {
	method string
	http.Handler
}*/

func NewRouter() *Router {
	root := node{component: "/", isNamedParam: false, methods: make(map[string]http.Handler)}
	return &Router{tree: &root}
}

func (r *Router) Add(method string, path string, h http.Handler) {
	if path[0] != '/' {
		log.Error(FormatError)
	}

	r.tree.addNode(method, path, h)
}

func (r *Router) Find(method string, path string, h http.Handler) {}

/*
func (r *Router) ShowHandler() map[string][]*Endpoint {
	return r.handlers
}*/

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	node, _ := r.tree.findNode(strings.Split(req.URL.Path, "/")[1:], req.Form)

	if handler := node.methods[req.Method]; handler != nil {
		handler.ServeHTTP(w, req)
	} else {
		if r.NotFound != nil {
			r.NotFound.ServeHTTP(w, req)
			return
		}
		http.Error(w, "URL Not Found", 404)
	}
}

func (r *Router) GET(path string, h http.Handler) {
	r.Add(GET, path, h)
}

func (r *Router) POST(path string, h http.Handler) {
	r.Add(POST, path, h)
}

func (r *Router) DELETE(path string, h http.Handler) {
	r.Add(DELETE, path, h)
}

func (r *Router) PUT(path string, h http.Handler) {
	r.Add(PUT, path, h)
}

func (r *Router) PATCH(path string, h http.Handler) {
	r.Add(PATCH, path, h)
}
