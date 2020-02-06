package main

import "fmt"

type person09 struct {
	id   int
	name string
	age  int
}

// 建议  使用指针类型
func (p *person09) PrintInfo() {
	fmt.Printf("编号%d\n", p.id)
	fmt.Printf("姓名%s\n", p.name)
	fmt.Printf("年纪%d\n", p.age)
}

//子类
type student09 struct {
	person09 //匿名字段
	class int
}

func (s *student09) PrintInfo() {
	fmt.Println("Student:", *s)
}
func main() {
	s := student09{person09{130, "小刚", 18}, 10}
	//子类对象方法 采用就近原则  使用子类的方法
	//方法重写
	s.PrintInfo()

	//父类对象方法
	s.person09.PrintInfo()
	fmt.Println(s.PrintInfo)
	fmt.Println(s.person09.PrintInfo)
}
