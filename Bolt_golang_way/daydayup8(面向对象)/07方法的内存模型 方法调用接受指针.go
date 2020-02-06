package main

import "fmt"

type student07 struct {
	name string
	age int
	score int
}
//对象不同 方法名相同 不会冲突
//在方法调用中  方法接收者是指针类型
//指针类型 普通类型 表示的是相同对象的类型【go给做的优化】
func (s *student07)Print()  {//调用该方法是 地址传递
	s.score = -9
	fmt.Println(*s)
}

func main() {
	stu := student07{"亚索",30,-5}
	//值传递 (s student07)   地址传递(s *student07)
	stu.Print()
	fmt.Println(stu)
}
