package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(r)
}