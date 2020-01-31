package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner2 interface {   //子集
	//方法  方法的声明
	sayhi()

}

type Personer interface {  //超集
	Humaner2   //继承sayhi()

	sing(string)
}

type student13 struct {
	name  string
	age   int
	score int
}

func (s *student13)sayhi()  {
	fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

func (s *student13)sing(name string)  {
	fmt.Println("我为大家唱首歌",name)
}

func main1301() {
	//接口类型变量定义
	var h Humaner2
	var stu student13 = student13{"小吴",18,59}
	h = &stu
	h.sayhi()

	//接口类型变量定义
	var p Personer
	p = &stu
	p.sayhi()
	p.sing("大碗面")
}

func main()  {
	//接口类型变量定义
	var h Humaner2  //子集
	var p Personer	//超集
	var stu student13 = student13{"小吴",18,59}

	p = &stu
	//将一个接口赋值给另一个接口
	//超集中包含所有子集的方法
	h = p  //ok

	h.sayhi()

	//子集不包含超集
	//不能将子集赋值给超集
	//p = h  //err
	//p.sayhi()
	//p.sing("大碗面")
}