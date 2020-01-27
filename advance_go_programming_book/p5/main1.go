package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("Hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timmeStart := time.Now()
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timmeStart)
		fmt.Print(timeElapsed)
	})
}
func main() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
}
