package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner1 interface {
	//方法  方法的声明
	sayhi()

}


type student12 struct {
	name  string
	age   int
	score int
}

func (s *student12)sayhi()  {
	fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

type teacher12 struct {
	name    string
	age     int
	subject string
}

func (t *teacher12)sayhi()  {
	fmt.Printf("大家好，我是%s,今年%d岁，我的学科是%s\n",t.name,t.age,t.subject)
}

//多态的实现
//将接口作为函数参数  实现多态
func sayhello(h Humaner1)  {
	h.sayhi()
}

func main() {

	stu := student12{"小明",18,98}
	//调用多态函数  不是将对象信息赋值给接口类型变量
	sayhello(&stu)

	tea := teacher12{"老王",28,"Go"}
	sayhello(&tea)
}
