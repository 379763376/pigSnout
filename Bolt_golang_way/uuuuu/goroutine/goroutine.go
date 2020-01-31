package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}(i)
	}
	//time.Millisecond一毫秒
	//time.Minute
	time.Sleep(time.Millisecond)
}

func main2()  {
	var a[10] int
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}