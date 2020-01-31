/*
变量的定义
cheat engine查看内存中的变量
变量是内存中的
内存地址的编号是16进制
可在函数内 或者直接放在包内，是一个包内变量，而非全局变量
*/
package main

import "fmt"

//函数外部定义变量,需要关键字var
var (
	aa2 = 3
	bb  = 4
)
var a, b, c bool

var d = 10
var (
	dd = 10
	aa = true
)

func main() {

	//一 、变量的声明
	//1.声明 var 变量名 类型
	//2.只是声明没有初始化 默认为0
	//3.在同一个{} 声明的变量要唯一
	var a int
	a = 10 //变量赋值
	fmt.Println(a)
	var b, c int
	b, c = 20, 30
	var s string
	println(b, c)
	fmt.Printf("%q\n", s) //会把quote打印出来

	//二 、变量的初始化  声明变量的时候赋值
	var d int = 10 //初始化 一步到位
	d = 20         //先要声明 再赋值
	println(d)

	var e float64 = 3.14
	println(e)

	var value float64 = 2
	//跟表达式
	var sum float64 = value * value
	println(sum)

	//三 、推倒类型
	//1.自动推倒
	var a1 = 10
	println(a1)
	fmt.Printf("%T\n", a1) //类型
	//2. 常用的自动推到
	//自动推倒  ：= 左边变量名没用过  只能用于第一次初始化
	b1 := 10
	var cc, dd = 1, "ddddd" //这种定义不是一种数据类型也可以写在一行自动推导
	fmt.Println(cc, dd)
	fmt.Println(b1)
	b1 = 20
	c1 := 3.14 //float64  不会是32
	fmt.Printf("%T\n", c1)
	x, y, z, ll := 10, 20, 3.33, "aaa"
	fmt.Println(x, y, z, ll)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", y)
	//其他
	//1.printf 格式化输出 printf("a = %d\n",a)  把a放到%d的位置
	//2.println 打印输出换行 printf("a = ",a)
	//3.print 打印输出 没有换行

	//四、多重赋值 和 匿名变量
	//多重赋值
	m, n := 10, 20
	fmt.Println(m, n)
	//1.数据交换
	//cpu分三层 计算+执行+寄存
	//寄存器存放数据  go直接在寄存器交换数据
	n, m = m, n
	fmt.Println(m, n)
	//2. _匿名变量  丢弃数据不处理
	temp, _ := 7, 8
	println(temp)
	//不需要第三个值
	var l, p, _ = test()
	println(l, p)

	//五、 多个变量或者常量的定义
	//变量 程序运行期间 可以改变的量 声明用var
	//1.不同类型的多个变量
	var (
		x1 int     = 1
		x2 float64 = 2.0
	)
	fmt.Println(x1, x2)
	//2.自动推倒类型
	var (
		x3 = 1
		x4 = 2.0
	)
	fmt.Println(x3, x4)
	x5, x6 := 1, 2.0
	fmt.Println(x5, x6)
	//常量 程序运行期间 不可以改变的量 声明用关键字 const
	//常量的定义
	const i int = 10
	const j float64 = 3.14
	fmt.Println(i, j)
	const (
		i2 int     = 10
		j2 float64 = 3.14
	)
	//自动推倒类型
	const (
		i3 = 10
		j3 = 3.14
	)
}
func test() (a, b, c int) {
	return 1, 2, 3
}
