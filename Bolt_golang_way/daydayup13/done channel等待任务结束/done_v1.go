package main

import (
	"fmt"
)

func doWork1(id int,w worker1) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)

		go func(){w.done <- true}()// 再开一个goroutine 并发发送done 就不会卡住
	}
}

type worker1 struct {
	in   chan int
	done chan bool
}

func createWorker1(id int) worker1 {
	w := worker1{
		in : make(chan  int),
		done :make(chan bool),
	}
	go doWork1(id,w)
	return w
}


func chanDemo1() {
	var workers [10]worker1
	for i := 0; i < 10; i++ {
		workers[i] = createWorker1(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	//for _,worker := range workers {
	//	<-worker.done
	//}


	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	//for _,worker := range workers {
	//	<-worker.done
	//}
	//wait for all of them
	for _,worker := range workers {
		<-worker.done
		<-worker.done
	}


}

func main() {
	chanDemo1()
}
/*
不要共享内存来通信 而是通过通信共享内存
用channel通知
 */