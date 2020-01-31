package main

import (
	"sync"
	"time"
)

//传统同步机制
/*
WaitGroup 用channel实现 但是里面还是同步机制
Mutex 互斥量
Cond
*/
//查看是否数据冲突：go run -race atimic.go
//尽量使用channel通信
import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex//互斥量
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	//atomic原子化的操作
	//atomic.AddInt32()  系统提供的atomic 多个goroutine执行是安全的
	//用Mutex互斥量实现atomic
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
