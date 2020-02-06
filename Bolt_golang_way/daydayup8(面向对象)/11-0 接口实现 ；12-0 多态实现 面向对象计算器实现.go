package main

import "fmt"

//定义接口
type Opter interface {
	//方法声明  因为需要返回 所以返回值int
	Result() int
}


//父类
type Operate struct {
	num1 int
	num2 int
}

//加法子类
type Add struct {
	Operate
}

//加法子类的方法
func (a *Add) Result() int {
	return a.num1 + a.num2
}

//减法子类
type Sub struct {
	Operate
}

//减法子类的方法
func (s *Sub) Result() int {
	return s.num1 - s.num2
}


//多态的实现
func Result(o Opter)  {
	v := o.Result()
	fmt.Println(v)
}

func main() {

	//1.通过对象方法调用
	//创建加法对象
	//var a Add
	//a.num1 = 10
	//a.num2 = 20
	//v := a.Result()
	//fmt.Println(v)

	//var s Sub
	//s.num1 = 10
	//s.num2 = 20
	//v := s.Result()
	//fmt.Println(v)

	//2.通过接口实现
	//var o Opter
	//var a Add = Add{Operate{10,20}}
	//o = &a
	//value := o.Result()
	//fmt.Println(value)

	//3.多态实现
	var a Add = Add{Operate{10,20}}
	Result(&a)
	var s Sub = Sub{Operate{10,20}}
	Result(&s)

}
