package main

import "fmt"

func main() {
	//1. 递归函数  调用函数本身
	//递归函数有两个相同结构：1。跳出条件
	test12(4)
	//2. 递归实现1+2+3+。。。+100
	//sum := add100(100)
	//fmt.Println(sum)
}
func test12(a int)  {
	if a == 10 {
		fmt.Println(a)
		return  //函数结束
	}
	test12(a+1)
	fmt.Println(a)
	//test12(4)[test12(5)[test12(6)[test12(7)[test12(8)[test12(9)[test12(10)]]]]]]
}
func add100(num int) int  {
	if num == 1{
		return 1
	}
	return num+add100(num - 1)
	//100+（99+（98+。。。+（3+（2+（1
}