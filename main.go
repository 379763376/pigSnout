package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ca := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i<10 ; i++{
		wg.Add(1)
		go func(i int) {
			time.Sleep(1000*time.Millisecond)
			defer wg.Done()
			ca<-i
		}(i)
	}

	go func() {
		wg.Wait()
		close(ca)
	}()

	for v := range ca{
		fmt.Println(v)
	}

}