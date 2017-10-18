package vidar

import (
	"net/http"
	"net/url"
	"path"
	"path/filepath"
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
	isStatic    bool
	handlers    map[string]http.Handler
}

func (n *node) addNode(method, path string, h http.Handler) {
	components := strings.Split(path, "/")[1:]
	c := components
	root := n

	for {
		oldNode, component := n.findNode(c, nil)

		if oldNode.component == component {
			oldNode.handlers[method] = h
			break
		}

		newNode := &node{
			height:    oldNode.height + 1,
			component: component,
			handlers:  make(map[string]http.Handler),
		}

		oldNode.children = append(oldNode.children, newNode)

		if len(component) > 2 && component[0] == ':' {
			newNode.isPathParam = true
		}

		// Ugly, Ugly, need to be refactored
		if component == "*" {
			root.handlers[method] = h
		}

		n = newNode
		c = components[newNode.height:]

		if len(c) == 0 {
			newNode.handlers[method] = h
			break
		}
	}
}

func (n *node) findNode(components []string, forms url.Values) (*node, string) {
	component := components[0]

	if len(n.children) > 0 {
		for _, child := range n.children {
			if component == child.component || child.isPathParam {
				if forms != nil && child.isPathParam {
					forms.Add(child.component[1:], component)
				}

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

//TODO: to be implemented
func (r *Router) Find(path string) {
	components := strings.Split(path, "/")[1:]
	r.tree.findNode(components, nil)
}

func (r *Router) Show() *node {
	return r.tree
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Form == nil {
		req.Form = make(url.Values)
	}

	node, _ := r.tree.findNode(strings.Split(req.URL.Path, "/")[1:], req.Form)
	log.Debug("We find %+v, which have handlers %+v", node.component, node.handlers)

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

func (r *Router) File(path, file string) {
	r.GET(path, ContextFunc(func(c *Context) {
		if err := c.File(file); err != nil {
			c.Error(err)
			c.Log.Error(err)
		}
	}))
}

func (r *Router) Static(prefix, root string, p *Plugin) {
	h := p.Apply(ContextFunc(func(c *Context) {
		if root == "" {
			root = "./"
		}

		if !strings.HasSuffix(root, "/") {
			root = root + "/"
		}

		upath := c.request.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
			c.request.URL.Path = upath
		}

		name := filepath.Join(root, path.Clean(upath)) // "/"+ for security
		if err := c.File(name); err != nil {
			c.Error(err)
			c.Log.Error(err)
		}
	}))

	r.GET(prefix, h)

	if prefix == "/" {
		r.GET(prefix+"*", h)
	}
}
