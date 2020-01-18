package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int,20)
	go func() {
		loop:
		for i :=0;;i++{
			time.Sleep(10*time.Millisecond)
			c1<-i
			fmt.Println("c1<-",i)
			if i == 30{
				close(c1)
				break loop
			}
		}
		fmt.Print("finish c1")
	}()
	loopM:
	for{
		time.Sleep(1000*time.Millisecond)
		select {
		case i,ok := <-c1:
			if !ok {
				break loopM
			}
			fmt.Println(i)
		}
	}
	fmt.Println("finish")

}