package main

import (
	"fmt"
	"time"
)

func main() {
	//for 后面不写任何东西  那么就是一个死循环 永远为真
	//break 跳出循环 如果多层for循环 跳出最近的内循环
	//continue 跳出本次循环 进入下一次循环
	//break 可以在for和switch使用，continue只能是在for
	//goto 跳转 goto关键字 End是标签 用户起的关键字
	// goto使用场景是测试的时候可以跳过一些不用注释代码

	i :=0
	for{
		i++
		time.Sleep(time.Second)
		if i == 5{
			break
		}
	}

	fmt.Println("aaa")
	goto End    //如果End在goto在goto之前了 就会无限制的循环
	fmt.Println("bbb")
	End:
	fmt.Println("ccc")
}
