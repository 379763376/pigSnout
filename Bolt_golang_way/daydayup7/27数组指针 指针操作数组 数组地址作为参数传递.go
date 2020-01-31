package main

import "fmt"

func main2101() {
	arr := [3]int{1, 2, 3}
	//定义指针  指向数组  数组指针
	var p *[3]int
	//指针和数组建立关系
	p = &arr
	fmt.Println(*p)

	//自动推导类型创建数组指针
	//p := &arr


	//通过指针操作数组
	//(*数组指针变量)[下标] = 值
	(*p)[0] = 111  //p是数组的地址 *p代表地址上的值==》代表的是数组的值  所以需要括起来  (*p)[0]指向数组的第一个元素
	//或者;
	p[1] = 222  //p

	fmt.Println(*p)

}

func main2102() {
	arr := [5]int{1, 2, 3, 4, 5}
	//指针变量和要存储数据类型要相同
	p1 := &arr
	p2 := &arr[0]
	//p1和p2在内存中指向相同的地址
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", p2)
	//指针的类型不同 一个数组指针 一个整型指针
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
}

func swaps(p *[5]int) {  //指针数组
	(*p)[0] = 111   //*p指向数组 [0]取出第一个元素
	p[1] = 222
}

func main() {
	//数组作为值传递 形参改变不会影响到实参
	//让形参改变能够影响实参 除了返回值 就需要用数组指针
	a := [5]int{1,2,3,4,5}
	swaps(&a)   //传地址 形参可以改变实参的值

	fmt.Println(a)

}
