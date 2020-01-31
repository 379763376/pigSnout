package main

import "fmt"

func main() {
	//1. 函数类型
	demo1(1,2)
	//2. 为什么函数名加一个括号就能代表函数
	//函数名接收一个地址 函数在内存区的地址
	fmt.Println(demo1) //输出的是一个低地址
	x := 10
	fmt.Println(&x) //输出的是一个高地址
	//==》函数名表示的是一个地址  函数在代码区的地址  （代码区  数据区  堆区 栈区)
	//内存四区模型 （从低地址到高地址   系统内部不做考虑--》（代码区  数据区  堆区 栈区)==》。。。  ）
	//代码区： 计算机指令 ：只读
	//数据区： 又分为：初始数据区、未初始化数据区、常量区
	//堆区：一个很大的空间 在使用时开辟空间，使用结束时释放
	//栈区: 函数内部定义的变量

	//3.用变量接收一个地址  通过f调用了函数
	f:=demo1
	f(1,2) //告诉我们这是代表一个函数
	fmt.Println(f) //地址和demo1函数的内存地址一样的
	fmt.Printf("%T\n",f)
	// 4.声明一个函数类型的变量
	var f1 func(int,int)
	fmt.Printf("%T\n",f1)
	//4. 定义函数类型 给已存在的函数起别名  （之前学过类型别名）
	var f2 FUNCDEMO1
	//fmt.Println(f2) //直接调用 没有赋值执行会报错  所以需要先给他赋值
	//要求demo1 的定义 和FUNCDEMO1要一致  包括：参数类型 个数 返回值
	f2 = demo1
	f2(10,20)
	fmt.Println(f2)
}
func demo1(a,b int)  {
	sum := a+b
	fmt.Println(sum)
}
type FUNCDEMO1 func(int, int)