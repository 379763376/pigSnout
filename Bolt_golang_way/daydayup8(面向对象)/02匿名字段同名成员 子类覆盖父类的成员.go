package main

import "fmt"

type Person2 struct {
	id   int
	name string
	age  int
}
//子类
type Student2 struct {
	Person2  //匿名字段
	name string
	score int
}
func main() {
	var s1 Student2

	//子类赋值
	//子类和父类结构体有相同的成员名 默认赋值给子类  采用就近原则
	s1.name = "花花"
	//父类赋值
	s1.Person2.name = "huahua"
	fmt.Println(s1)
}
