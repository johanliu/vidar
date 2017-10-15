package vidar

import (
	"errors"
	"net/http"
)

type Ring func(http.Handler) http.Handler

type Plugin struct {
	rings []Ring
}

type ContextFunc func(*Context)

var errNilHandler = errors.New("handler: nil handler can't be used")

// ServeHTTP calls f(w, r).
func (f ContextFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	c.Log = log
	f(c)
}

func NewPlugin(ring ...Ring) *Plugin {
	return &Plugin{append([]Ring{}, ring...)}
}

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

func (p *Plugin) Append(ring Ring) {
	if ring == nil {
		log.Error(errNilHandler)
	}

	p.rings = append(p.rings, ring)
}

func (p *Plugin) Extend(np Plugin) Plugin {
	newRings := make([]Ring, len(p.rings)+len(np.rings))
	copy(newRings, p.rings)
	copy(newRings, np.rings)
	return Plugin{newRings}
}
