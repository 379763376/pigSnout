package main

import (
	"fmt"
	"os"
	"sync"
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
type ByteCounter int

func (c *ByteCounter)Write(p []byte) (n int, err error)  {
	*c += ByteCounter(len(p))
	return len(p),nil
}
func (c *ByteCounter)String(){
	fmt.Fprint(os.Stdout,c)
}

func toPrint(stringer *ByteCounter, s string) {
	stringer.String()
}
func main() {
	var c  ByteCounter
	fmt.Fprintf(&c,"","")
	toPrint(&c,"xxxx")
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
