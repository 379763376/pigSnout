package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {

	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}
func worker_old(id int,c chan int)  {
	//n,ok := <-c 可以用ok判断是否是close
	//或者range
	for {
		n,ok := <-c
		if !ok{
			fmt.Println("aaa")
			break
		}
		fmt.Printf("worker %d received %d\n",
			id,n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)//创建了channel
	go worker(id, c) //实际的协程工作内容,提出来作为共用
	return c
}

func chanDemo() {
	var channels [10]chan<- int //定义了10个channel接收数据
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) //缓存大小为3
	go worker_old(0, c) //
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'//如果没有接收channel 会deadlock 死锁
	time.Sleep(time.Millisecond)
}

func channelClose() { //channel如果有明确结尾 发送发通知接收方自己没有数据要发了
	c := make(chan int)
	go worker_old(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) //即便关闭
	time.Sleep(time.Second)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
/*
channel
用作两个routine之间数据的交换
var c chan int //定义一个channel 传入的数据是int，但是这个channel没给做出来是c == nil
所以创建一个channel需要c:make(chan int)

chanDemo中定义一个channel然后定义一个routine接收数据
不接收会datalock异常，
c<- 发数据到channel
n := <-c 从channel取出数据

函数是一等公民 channel也是 可以作为函数 可以作为参数 可以作为返回值

bufferChannel 添加缓冲区

channel close 发送发close，外层函数推掉 channel也就关闭了

channel的理论基础CSP communication sequential process

CSP：
不要通过共享内存来通信；通过通信共享内存
chanDemo发完数据后work通知打印完毕了
 */