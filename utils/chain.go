package utils

import "net/http"

type Ring func(http.Handler) http.Handler

type Chain struct {
	rings []Ring
}

func New(ring ...Ring) Chain {
	return Chain{append([]Ring{}, ring...)}
}

func (c Chain) Wrap(h http.Handler) http.Handler {
	for i := range c.rings {
		h = c.rings[len(c.rings)-(i+1)](h)
	}
	return h
}

func (c Chain) Append(ring Ring) Chain {
	return Chain{append(c.rings, ring)}
}

func (c Chain) Extend(newChain Chain) Chain {
	newRings := make([]Ring, len(c.rings)+len(newChain.rings))
	copy(newRings, c.rings)
	copy(newRings, newChain.rings)
	return Chain{newRings}
}
