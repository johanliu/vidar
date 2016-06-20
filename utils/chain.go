package utils

import (
	"net/http"

	"github.com/johanliu/Vidar/context"
	"github.com/johanliu/Vidar/logger"
)

type Ring func(http.Handler) http.Handler

type Chain struct {
	rings []Ring
}

type ContextWrapperFunc func(*context.Context)

// ServeHTTP calls f(w, r).
func (f ContextWrapperFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := context.New(w, r)
	f(c)
}

func New(ring ...Ring) Chain {
	return Chain{append([]Ring{}, ring...)}
}

func (c *Chain) Wrap(f ContextWrapperFunc) http.Handler {
	return c.wrapInternal(ContextWrapperFunc(f))
}

func (c *Chain) wrapInternal(h http.Handler) http.Handler {
	if h == nil {
		logger.Error.Println("Handler can not be nil to be wrapped")
	}

	for i := range c.rings {
		h = c.rings[len(c.rings)-(i+1)](h)
	}
	return h
}

func (c *Chain) Append(ring Ring) {
	if ring == nil {
		logger.Error.Println("nil Handler can not be appended ")
	}

	c = &Chain{append(c.rings, ring)}
}

func (c *Chain) Extend(newChain Chain) Chain {
	newRings := make([]Ring, len(c.rings)+len(newChain.rings))
	copy(newRings, c.rings)
	copy(newRings, newChain.rings)
	return Chain{newRings}
}
