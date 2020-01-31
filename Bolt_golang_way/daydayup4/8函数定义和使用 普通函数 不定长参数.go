package main

import (
	"fmt"
)

func main() {
	//1.函数  定义：将代码重用的机制
	//2 函数的基本语法： func 函数名（参数列表） 返回值 {...}
	// 函数不能重名
	//3 函数的使用 run()
	fmt.Println(len("hello"))
	run()
	//4 普通参数
	test("aa") //实参  编译器的参数名提示功能
	test2(1,2)
	test3(1,"1")
	//3 不定长参数
	test4(1,2,3,4)
	//4. 不定参要放到最后，不能在其他定参前面
	test5("a",1,2,3,4)
	//固定参数必须传值，不定参数根据需要决定是否传值
}
func run()  {
	fmt.Println("run")
}
func test(a string)  { //形参
	fmt.Println(a)
}
func test2(a1,b1 int)  {
	fmt.Println(a1+b1)
}
func test3(a1 int,b1 string)  {
	fmt.Println(a1,b1)
}
func test4(t ...int) {//不定参  在调用函数时传不定数量的参数 通过下标取值
	fmt.Println(t)
	sum := 0
	for i:=0;i<len(t) ; i++ {
		sum +=t[i]
	}
	fmt.Println(sum)


	//或
	sum2 :=0
	for _,v := range t {
		sum2+=v
	}
	fmt.Println(sum2)
}
func test5(x1 string,x2 ...int) {
	fmt.Println(x1,x2[0],x2[1],x2[2],x2[3])
}