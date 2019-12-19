package main

import (
	"sync"
	"time"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	//mapping: make(map[string]string),
	mapping: map[string]string{
		"zs": "lz",
	},
}

func main() {
	Lookup("zs")
	time.AfterFunc(10*time.Second, func() { r.Launch() })
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
