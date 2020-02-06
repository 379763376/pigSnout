package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriver2 struct {
	UserAgent string
	TimeOut time.Duration
}

func (r *Retriver2) Get(url string) string {
	resp,err := http.Get(url)
	// Caller should close resp.Body when done reading from it.
	if err != nil {
		panic(err)
	}
	result,err := httputil.DumpResponse(resp,true)

	resp.Body.Close() //暂时先这么写

	if err !=nil {
		panic(err)
	}
	return string(result)
}

