package main

import "fmt"


type A struct {
	a int
}
type B struct {
	A
	b int
}
type I interface {
	say()
	hello()
	hi()
}
func (a *A)say()  {
	fmt.Println("a")
}

func (a *A)hello()  {
	fmt.Println("aa")
}

func (a *A)hi()  {
	fmt.Println("aaa")
}
func main() {
	var b B
	(&b).hi()
	b.hi()
	var i I = &b
	i.hi()
}