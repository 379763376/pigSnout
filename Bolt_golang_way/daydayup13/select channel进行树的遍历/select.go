package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)//time.Duration一段时间
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]

		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println(
				"queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
//1.创建管道，往管道发送数据
func main1() {
	var c1,c2 = generator(),generator()
	for{
		select {
		case n := <-c1:
			fmt.Println("c1",n)
		case n:=<-c2:
			fmt.Println("c1",n)
		default:
			fmt.Println("no")

		}
	}
}
//
func main3() {
	var c1,c2 = generator(),generator()
	var worker = createWorker(0) //初始化的channnel
	n := 0
	hasValue := false
	for{
		var activeWorker chan<- int //activeWorker == nil
		if hasValue{
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		}
	}
}
func main4() {
	//生成数据的速度大于消耗速度，这一段时间内的数据会被冲掉，解决方式：把收到的存下来排队
	//希望程序运行10s停止
	//如果800ms没有数据 提示time-out
	//定时查看队列的积压 使用定时器
	var c1,c2 = generator(),generator()
	var worker = createWorker(0) //初始化的channnel
	var values []int

	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for{
		var activeWorker chan<- int
		var activeValue int

		if len(values)>0{
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values,n)
		case n := <-c2:
			values = append(values,n)
		case activeWorker <- activeValue:
			values = values[1:]

		case <-time.After(800 * time.Millisecond)://是每次select花的时间，两次生成数据时间差超过800
			fmt.Println("timeout")
		case <-tick: //定时器
			fmt.Println(
				"queue len =", len(values))

		case <-tm://从这个channel获取数据,是总的时间
			fmt.Println("bye")
			return
		}
	}
}
