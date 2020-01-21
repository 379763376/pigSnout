package main

import "fmt"

type IS interface {
	ISPrint()
}
func CallISprint(is IS){
	is.ISPrint()
}
type ISIMPL struct {
}

func (is ISIMPL)ISPrint()  {
	fmt.Println("starting")
}
func main() {
	var is IS = new(ISIMPL)
	CallISprint(is)
}
