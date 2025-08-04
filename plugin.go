package vidar

import (
	"errors"
	"net/http"
)

// Ring represents a middleware function that wraps an HTTP handler.
type Ring func(http.Handler) http.Handler

// Plugin manages a collection of middleware (rings) for HTTP request processing.
type Plugin struct {
	rings []Ring // Ordered list of middleware functions
}

// ContextFunc is a handler function that takes a Vidar Context instead of standard HTTP parameters.
type ContextFunc func(*Context)

var errNilHandler = errors.New("handler: nil handler can't be used")

// ServeHTTP implements http.Handler for ContextFunc, creating a Context and calling the function.
func (f ContextFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	c.Log = log
	f(c)
}

// NewPlugin creates a new Plugin instance with optional initial middleware.
func NewPlugin(ring ...Ring) *Plugin {
	return &Plugin{append([]Ring{}, ring...)}
}

// Apply wraps a ContextFunc with all registered middleware in the plugin.
func (p *Plugin) Apply(f ContextFunc) http.Handler {
	return p.useInternal(ContextFunc(f))
}

func (p *Plugin) useInternal(h http.Handler) http.Handler {
	if h == nil {
		log.Error(errNilHandler)
	}

	for i := range p.rings {
		h = p.rings[len(p.rings)-(i+1)](h)
	}
	return h
}

// Append adds a middleware function to the end of the plugin's middleware chain.
func (p *Plugin) Append(ring Ring) {
	if ring == nil {
		log.Error(errNilHandler)
	}

	p.rings = append(p.rings, ring)
}

// Extend creates a new Plugin by combining the current plugin with another plugin's middleware.
func (p *Plugin) Extend(np Plugin) Plugin {
	newRings := make([]Ring, len(p.rings)+len(np.rings))
	copy(newRings, p.rings)
	copy(newRings[len(p.rings):], np.rings)
	return Plugin{newRings}
}
