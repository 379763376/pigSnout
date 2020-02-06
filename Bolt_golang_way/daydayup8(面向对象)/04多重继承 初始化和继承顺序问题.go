package main

import "fmt"

//父类
type person04 struct {
	name string
	age  int
	sex  string
}

//父类
type Person04 struct {
	id   int
	addr string
}

//子类
type Student04 struct {
	//结构体成员为多个匿名字段
	Person04 //匿名字段
	person04
	score int
}

func main() {
	var stu Student04
	stu.id = 200
	stu.addr = "北京"
	stu.name = "王富贵"
	stu.age = 10
	stu.sex = "男"
	stu.score = 100
	fmt.Println(stu)

	//自动推导类型
	s := Student04{Person04{200,"北京"},person04{"王富贵",10,"男"},100}
	fmt.Println(s)
}
