package mock

import "fmt"

type Retriver1 struct {
	Contents string
}

func (r *Retriver1) String() string {
	return fmt.Sprintf("Retriver: {Contents=%s}",r.Contents)
} //实现了fmt中的stringer 类似与java的toString

func (r *Retriver1) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}


func (r Retriver1) Get(url string) string {
	return r.Contents
}


