package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片的地址和扩容机制


	//一.1.地址：
	//空切片 指向内存地址 编号为0的空间
	//切片名本身就是一个地址
	var slice [] int
	fmt.Printf("%p\n",slice) //输出0x0  直接使用切片名  不用&  因为它本身就是一个地址

	//追加数据
	//当使用append追加数据时 切换地址可能会发生改变
	slice = append(slice,1,2,3)
	fmt.Printf("%p\n",slice) //输出0xc000060140
	slice = append(slice,4,5,6)
	fmt.Printf("%p\n",slice) //输出0xc00007a100
	/*
	切片名放在栈区，值为内存地址
	切片中的数据放在堆区

	堆区和栈区存放数据的差别
	栈区要比堆区空间小的多
	栈区存放已经知道数据类型和大小的数据
	堆区存放大小不确定的（切片可以追加数据)
	 */
	var slice2 [] int
	fmt.Printf("%p\n",slice2) //输出0x0
	//打印第一个元素的地址
	fmt.Printf("%p\n",&slice[0])//输出0xc00007a100 和数组一样 内存地址以数组下标为0的位置作为切片的地址
	fmt.Printf("%p\n",&slice[1])//输出0xc00007a108
	//append的时候数据地址为什么会发生改变？？？
	//当追加数据的时候 后面没有可申请容量  只能是另找可以满足容量大小的内存存储  这样就会发生改变




	//二.扩容
	//容量为1
	s := make([] int,0,1)
	fmt.Println(s)
	//扩容 初期都是在原始数据的基础上扩大两倍，当容量增加到1024以后  扩容上次cap的1/4
	//每次扩容都是2的倍数
	//for i := 0; i< 100000 ;i++  {
	//	oldcap := cap(s)
	//	s = append(s, i)
	//	newcap := cap(s)
	//	if(newcap>oldcap){
	//		fmt.Println(oldcap,newcap)
	//	}
	//}

	//三 切片操作
	//1.截取
	// s[起始位置：结束位置(不包含该位置)：计算容量的数据]  [low:high:max]
	a3 := []int{10,20,30,40,50} //cap = 5
	//len = high - low
	//cap = max - low  max不能超过原始数据切片的容量5
	//b3 := a3[0:4:6]
	fmt.Println(cap(a3))
	b3 := a3[1:4:5]
 	//a3[:]等价于:= a3
	//a3[:3]等价于 [0:4]
	//a3[0:]等价于:= a3
	//[2:3]
	fmt.Println(b3)
	//对截取出来的数据修改
	b3[0] = 0
	//打印发现a3发生改变 截取后的切片仍是原来切片中的内容
	fmt.Println(a3)
	fmt.Printf("%p\n",a3)
	fmt.Printf("%p\n",&a3[1])
	fmt.Printf("%p\n",b3)

	//2.切片的内置函数
	//append
	//copy  在内存中存储两个独立的切片内存  如果任意一个发生修改 不会影响到另一个
	var c3 [] int = [] int {1,2,3,4,5}
	d3 := make([]int ,6)
	copy(d3,c3)
	fmt.Println(d3)
	fmt.Printf("%p\n",c3)
	fmt.Printf("%p\n",d3)
	d3[4]=6
	fmt.Println(d3)
	c3 = append(c3, 6,7,8,9)
	fmt.Println(c3)
	e3 := []int{3,2,1}
	copy(e3,c3)
	fmt.Println(e3) //结果时1，2，3因为长度为3 所以拷贝的时候只会从c3中拷贝两个
	f3 := []int{9,8,7,6,5,4,3,2,1}
	copy(f3,e3)
	fmt.Println(f3) //结果为1，2，3，6，5，4，3，2，1  覆盖了前三位


	//四.
	// 排序  冒泡排序
	sort.Ints(f3)
	//数组作为函数的参数 是值传递
	//实参和形参 是不同的地址 形参不能改变实参值
	//可以通过返回值改变实参的值
	//和数组作为参数传递的区别：
	f4 := []int{1,2,3,4,5}
	foo(f4)
	fmt.Println(f4) //实参也改变了   切片作为函数的参数传递的过程是地址传递【是引用传递】形参可以改变实参的值
	f4 = foo(f4)
	fmt.Println(f4) //对切片追加的时候 切片地址可能会发生改变，形参的地址空间改变了 指向的地址空间变了，所以值不一样
	//返回值
	//返回值是修改了实参的地址 指向了新的地址所以也就改了
	//切片在内存中存储 如果该空间没有指向 进行销毁

}

func foo(f4 []int)  [] int{
	f4[0] = 2
	//fmt.Println(f4)
	f4 = append(f4, 6,7,8,9,10)
	return f4
}