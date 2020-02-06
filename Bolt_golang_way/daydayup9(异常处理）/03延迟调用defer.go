package main

import "fmt"

func test03(x int) {
	v := 100 / x
	fmt.Println(v)
}
func main() {
	//defer语句只能出现在函数内部
	//defer fmt.Println("hello")
	//defer fmt.Println("老王")
	//defer fmt.Println("你好")
	//defer执行顺序
	//一个函数有多个defer语句  它们先进后出的顺序执行

	//基石函数或者某个延迟调用发生错误 这些调用依旧会被执行
	//defer fmt.Println("hello")
	//defer fmt.Println("老王")
	//defer test03(0)
	//defer fmt.Println("你好")

	a := 10
	b := 20
	defer func(a,b int) {  //传参后值发生改变。。。虽然程序没有立即执行，但是已经完成了值的传递 a=10,b=20
		fmt.Println("匿名函数a", a)
		fmt.Println("匿名函数b", b)
	}(a,b)

	a = 100
	b = 200
	fmt.Println("main函数a", a)
	fmt.Println("main函数b", b)
}
