package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner interface {
	//方法  方法的声明
	sayhi()

}


type student11 struct {
	name  string
	age   int
	score int
}

func (s *student11)sayhi()  {
	fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

type teacher11 struct {
	name    string
	age     int
	subject string
}

func (t *teacher11)sayhi()  {
	fmt.Printf("大家好，我是%s,今年%d岁，我的学科是%s\n",t.name,t.age,t.subject)
}




func main() {
	//接口是一种数据类型 可以接收满足对象的信息
	//接口是虚的  方法是实的
	//接口定义规则  方法实现规则
	//接口定义的规则  在方法中必须有定义的实现
	var h Humaner

	stu := student11{"小明",18,98}
	//stu.sayhi()
	//将对象信息赋值给接口类型变量
	h = &stu
	fmt.Println(h)
	h.sayhi()

	tea := teacher11{"老王",28,"物理"}
	//tea.sayhi()
	//将对象赋值给接口 必须满足接口中的方法的声明格式
	h = &tea
	h.sayhi()
}
