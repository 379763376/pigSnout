package main

import (
	"fmt"
)

func main() {
	timeout := [5]string{"1.111111111 1111","10","100","1000","10000"}
	for _,v := range timeout{
		fmt.Printf("%.5s\n",v)
	}

}
