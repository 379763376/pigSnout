package main

import "fmt"

func main() {
	//基本数据int作为参数
	a :=1
	b :=2
	f(a,b)
	//==>以上a,b的结果没有发生变化
	//1. 数组作为参数传递
	a1 := [2]int{1,2}
	f1(a1)
	fmt.Println(a1)
	//==>数组作为函数参数是值传递，而且实参和形参是不同的地址，内存中是有两份独立的数组 存储不同的数据
	//在函数调用结束后 内存回收 不会影响到主函数实参的值
	//==》如果希望函数计算结果传递给实参，函数需要返回值
	a1 = ff1(a1)
	fmt.Println(a1)

	//2. 二维数组 [行][列]
	var a2 [2][3]int
	fmt.Println(a2)
	//len(a2) 行数
	//len(a2[1]) 或者 len(a2[0]) 列数 都是3
	//赋值
	a2[0][0] = 0
	a2[0][1]=1
	a2[0][2]=2
	a2[1][0]=3
	a2[1][1]=4
	a2[1][2]=5
	fmt.Println(a2)
	//外层控制行  内层控制列
	for i:=0;i<len(a2) ;i++  {
		for j:=0;j<len(a2[0]) ;j++  {
			fmt.Println(a2[i][j])
		}
	}
	//3. 二维数组初始化
	//全部初始化 部分初始化
	a3 := [3][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Println(a3)
	/*
	4.数组的长度是固定的
	5.数组作为函数参数传递 定义函数形参的长度需要指定长度
	*/
	/*
	6.切片
	可以不固定长度，可以追加元素
	 */
	//定义方式一：不初始化
	var slice [] int //空切片
	fmt.Println(slice)
	slice=append(slice, 1,2,3,4,5)
	fmt.Println(slice)
	fmt.Println(len(slice)) //长度
	fmt.Println(cap(slice)) //容量  可以申请预留后面的若干位置
	//定义方式二：初始化
	var a4 []int=[]int{1,2,3}
	fmt.Println(a4)
	a4=append(a4,4,5,6)
	//自动推导：
	b4 :=[]int{1,2,3}
	fmt.Println(b4)
	//定义方式三：[用的不多] 通过make函数实现的
	c4 := make([]int,5,10) //类型 长度5 容量10。定义了一个空切片，长度为5，容量为10
	fmt.Println(c4,len(c4),cap(c4))
	c4 = append(c4, 1,2,3,4,5,6)
	//后面这些append追加数据放到容量里，而不是初始化的5个
	//[00000123456]
	//长度小于容量 ，注意容量需要大于长度，
	//c4 := make([]int,5,5) 会报错
	//长度为已经初始化的空间，容量是已经开辟的空间 预留出来的

	//容量可以不写省略掉 切片根据添加的数字自己添加
	d4 := make([]int,5)
	fmt.Println(d4)
	//赋值 下标方式 循环方式
	d4[0] = 1
	d4[1]= 2
	for i := 0;i<len(d4) ;i++  {
		d4[i] = i
	}
	//注意循环赋值使用的是len(d4）而不能是cap 会报越界，不能使用容量做边界条件



}
func f(a,b int)  {
	a,b = b,a
}
func f1(arr [2]int)  {
	arr[0],arr[1] = arr[1],arr[0]
}
func ff1(arr [2] int) [2]int  {
	arr[0],arr[1] = arr[1],arr[0]
	return arr
}

