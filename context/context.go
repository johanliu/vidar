package context

import (
	"net/http"
	"sync"
)

var (
	mutex sync.RWMutex
	ctx   = make(map[*http.Request]map[string]interface{})
)

//TODO
func Set(r *http.Request, key string, value interface{}) {}

func Get(r *http.Request, key string) {}
