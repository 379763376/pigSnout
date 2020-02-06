package main

import "fmt"

type cat struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

//函数定义

//func 函数名 (参数列表) 返回值列表 {
//	代码体
//}

func show()  {
	fmt.Println("喵喵叫")
}
//方法定义
//func (对象)方法 (参数列表) 返回值列表 {
//	代码体
//}


//方法  绑定对象的

//结构体类型 可以作为对象类型
//结构体作为接收者
func (c cat)show()  {
	fmt.Println("喵喵叫")
	fmt.Printf("我是%s,喵喵叫\n",c.name)
}

//方法名一样 接受者不一样 方法也就不一样
func (d dog)show()  {
	fmt.Println("汪汪叫")
}


func main() {
	//对象创建
	var c cat

	c.name = "小花"
	c.age = 2
	fmt.Println(c)
	//对象.方法   包.函数  结构体.成员
	c.show()

	//函数调用
	//show()


	var d dog
	d.name = "旺财"
	d.age = 3
	fmt.Println(d)
	d.show()


	//对象的方法名和函数名可以 重名  但是 相同的对象方法不能重名
	//show()


}
