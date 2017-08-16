package vidar

import (
	"net/http"
)

type Ring func(http.Handler) http.Handler

type Chain struct {
	rings []Ring
}

type ContextUserFunc func(*Context)

// ServeHTTP calls f(w, r).
func (f ContextUserFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	f(c)
}

func NewChain(ring ...Ring) Chain {
	return Chain{append([]Ring{}, ring...)}
}

func (c *Chain) Use(f ContextUserFunc) http.Handler {
	return c.useInternal(ContextUserFunc(f))
}

func (c *Chain) useInternal(h http.Handler) http.Handler {
	if h == nil {
		Error.Println("Handler can not be nil to be wrapped")
	}

	for i := range c.rings {
		h = c.rings[len(c.rings)-(i+1)](h)
	}
	return h
}

func (c *Chain) Append(ring Ring) {
	if ring == nil {
		Error.Println("nil Handler can not be appended ")
	}

	c.rings = append(c.rings, ring)
}

func (c *Chain) Extend(newChain Chain) Chain {
	newRings := make([]Ring, len(c.rings)+len(newChain.rings))
	copy(newRings, c.rings)
	copy(newRings, newChain.rings)
	return Chain{newRings}
}
