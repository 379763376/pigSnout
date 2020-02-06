package main

import "fmt"

func test()  {
	//定义在函数内部的变量 当函数销毁的时候 变量也会跟随消失
	a :=10 //局部变量  作用于函数内部
	fmt.Println(a)
	b = 30
	fmt.Printf("test:%d\n",b)
}
var b int
 var c = 20
 var d int
func b1()  {
	fmt.Println(b)
}
func main() {
	//1. 函数名在当前包下唯一
	//2. 变量的作用域
	a := 1  //局部变量 作用域 函数内部有效
	//test()
	fmt.Println(a)
	//3. 变量先声明  再使用 在函数内部变量时唯一的
	//4. for循环  中定义的初始化变量可以和函数内部变量同名？  可以把for循环理解为一个函数
	//5. 全局变量 定义在函数外部的变量  在整个包里要唯一，不能和其他文件的全局变量重名
	//       全局变量名和局部变量名可以重名 尽量避免
	b+=1
	b++
	fmt.Println(b)
	b = 20
	test()
	fmt.Printf("main:%d\n",b)
	//6. 全局和局部变量的关系 重名会使用局部变量 就近原则
	b := 100
	fmt.Printf("main:%d\n",b)
	//7.全局变量在函数中被修改  修改局部变量 会影响到其他位置使用全局变量
	b1()

	//
	fmt.Printf("%v b1函数 代码区地址\n",b1) //代码区
	fmt.Printf("%v c数据区 全局变量地址 初始化\n",&c) //初始化数据区
	fmt.Printf("%v d数据区 全局变量地址 未初始化\n",&d) //未初始化数据区
	b2 := 1
	fmt.Printf("%v b2栈区 局部变量地址\n",&b2) //栈区

}


