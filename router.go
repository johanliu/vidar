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
	height      int
	children    []*node
	component   string
	isPathParam bool
	handlers    map[string]http.Handler
}

func (n *node) addNode(method, path string, h http.Handler) {
	components := strings.Split(path, "/")[1:]
	c := components

	for {
		current, component := n.findNode(c, nil)

		if current.component == component {
			current.handlers[method] = h
			break
		}

		newNode := &node{
			height:    current.height + 1,
			component: component,
			handlers:  map[string]http.Handler{method: h},
		}

		if len(component) > 2 && component[0] == ':' {
			newNode.isPathParam = true
		}

		current.children = append(current.children, newNode)

		n = newNode
		c = components[newNode.height-1:]
	}
}

func (n *node) findNode(components []string, forms url.Values) (*node, string) {
	component := components[0]

	if len(n.children) > 0 {
		for _, child := range n.children {
			if child.isPathParam && forms != nil {
				forms.Add(child.component[1:], component)
			}

			if component == child.component {
				next := components[1:]
				if len(next) > 0 {
					return child.findNode(next, forms)
				} else {
					return child, component
				}
			}
		}
	}

	return n, component
}

func NewRouter() *Router {
	root := node{height: 0, component: "/", isPathParam: false, handlers: make(map[string]http.Handler)}
	return &Router{tree: &root}
}

func (r *Router) Add(method string, path string, h http.Handler) {
	if path[0] != '/' {
		log.Error(FormatError)
	}

	r.tree.addNode(method, path, h)
}

func (r *Router) Find(method string, path string, h http.Handler) {
	components := strings.Split(path, "/")[1:]
	r.tree.findNode(components, nil)
}

func (r *Router) Show() *node {
	return r.tree
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	node, _ := r.tree.findNode(strings.Split(req.URL.Path, "/")[1:], req.Form)

	if handler := node.handlers[req.Method]; handler != nil {
		handler.ServeHTTP(w, req)
	} else {
		//TODO: error handler
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
