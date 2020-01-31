package main

import (
	"fmt"
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

	time.Sleep(time.Minute)
}
/*

以下直接引用for中的i不安全 所以需要上面方法传进去

for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}()
	}
for i := 0; i < 10; i++ {
		go func() {
			for {
				a[i]++
				runtime.GoSched()
			}
		}()
	}
go run 1.goroutine.go
go run -race goruntine.go 检测数据访问冲突
main.for执行完后i编程10，再执行go协程的时候a[10]是越界了


开了10个goroutine协程
main和协程是并发执行的，主程序执行完了，goroutine也会被杀掉了 time.Millisecond毫秒
go并发执行开1000个协程 coroutine

coroutine是轻量级“线程”
线程在任何时候都有可能被系统切换 是抢占式任务处理，系统处理完其他的还会回来继续执行完没完成的线程  需要存更多的上下文关系
非抢占式多任务处理 由协程主动交出控制权 不同于抢占式，要保存各线程的状态 只用保存资源切换的几个点
编译器/解释器/虚拟机层面的多任务   编译器级别的多任务 go有调度器调度协程 操作系统也有调度器调度协程
多个协程可能在一个或多个线程上运行
创建了一个10长度的数组，每个数组创建一个协程 分别对数据增加
a[i]++ 是不会主动交出控制权，可以IO交出控制权，也可以手动交出runtime.GoSched()  让别人有机会执行  每个人获得的机会是不一样的

子程序是协程的一个特例
协程和普通函数区别
普通函数 在一个线程内  main调用函数，等执行完再执行下一个语句
协调 在一个或多个线程内 main和dowork之间数据和控制权可以互相通信，交换控制权

其他语言对协程的支持
c++ boost.coroutine
python 3.5 asnyc def对协程原生支持，之前是使用yield关键字实现协程

调度器控制是否放一个线程里

goroutine的定义
任何函数只需要加上go就能送给调度器运行
不需要在定义时区分是否时异步函数
调度器在合适的点进行切换【切换的点并不能完全的控制】
使用-race检测数据访问的冲突

goroutine可能的切换点[只是参考 不能保证切换 不能保证在其他地方不切换]
i/o ,select channel进行树的遍历
channel
等待锁
函数调用（调度器决定是否切换）
runtime.Gosched()

查看开启一千个协程 启动了多少线程
查看实际系统上只是最多4个线程  因为4个核

*/