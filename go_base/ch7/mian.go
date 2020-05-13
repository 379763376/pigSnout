package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer //接口
	w = os.Stdout //w为*File类型,具体实现类型
	f := w.(*os.File) //f接口值 T:os.Stdout V:
	
	fmt.Println(f)
	c := w.(*bytes.Buffer)
	fmt.Println(c)
}