package main

import "fmt"

type Person3 struct {
	id   int
	name string
	age  int
}
//子类
type Student3 struct {
	*Person3  //   指针类型匿名字段
	score int
}
func main() {

	var s Student3
	s.score = 90
	s.Person3 = new(Person3)  //创建一块空间 指针指向一个有效的地址
	s.name = "李宁"

	//(*s.Person3).name = "周杰伦"
	s.id = 110
	s.age = 20
	fmt.Println(s.name,s.id,s.age) //打印时不能直接打印s  需要s.id

	s.Person3 = &Person3{101,"林俊杰",35}
	fmt.Println(s.name,s.id,s.age)
	
}
