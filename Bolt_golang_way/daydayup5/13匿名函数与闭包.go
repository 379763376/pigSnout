package main

import "fmt"

func main() {
	//匿名函数与闭包
	//一, 匿名函数
	// 没有名字的函数 在函数中定义一个函数
	var num int
	num = 9
	f := func() {  //匿名函数需要一个变量去接收
		num ++  //在匿名函数里面是可以直接使用主函数中的局部变量   匿名函数对主函数的修改  影响到主函数的值
		fmt.Println(num)
	}
	//2.函数的调用
	f()  //方法一 函数的调用
	fmt.Println(num)
	func() {
		num ++
		fmt.Println(num)
	}() //方法二 函数的调用
	//方法三 函数的调用  自己定义一个函数类型 创建一个变量 调用函数
	type FType func()
	var func1 FType
	func1 = f
	func1()
	//3. 有参数传递  有返回值的匿名函数
	func(a,b int) {
		sum :=a+b
		fmt.Println(sum)
	}(3,6)
	x,y,z := func(a,b int)(i,j,sum int){
		i= a
		j =b
		sum = a+b
		return

	}(1,3)
	fmt.Println(x,y,z)
	// 二 闭包
	//1. 匿名函数 实现了闭包 ： 在一个函数中有权访问另外一个函数作用域中变量的函数  匿名函数访问main中定义的num
	// 以上关于闭包函数的定义 可见 匿名函数就是闭包
	fmt.Println(test1())
	fmt.Println(test1())  //返回值始终是1 因为上次调用完后就销毁了
	f11:= test2()  //使用变量接收函数
	//闭关不关心捕获的变量 或者常量是否超出了作用域  只要闭包还在使用它 这个变量就还存在 所以就不会重复的初始化x的操作
	fmt.Println(f11())  //f11是地址 加括号是函数的调用
	fmt.Println(f11())  //输出的是 2， 数字变成累加
}
func test1() int {
	var x int
	x++
	return x
}
func test2() func() int{ //返回匿名函数类型
	var x int
	return func() int{  //返回匿名函数
		x++
		return x
	}
}