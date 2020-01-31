package main

import "fmt"

//父类
type person08 struct {
	id int
	name string
	age int
}

//子类
type student08 struct {
	person08   //匿名字段
	class int
}
// 建议  使用指针类型
func (p *person08)PrintInfo()  {
	fmt.Printf("编号%d\n",p.id)
	fmt.Printf("姓名%s\n",p.name)
	fmt.Printf("年纪%d\n",p.age)
}

func main() {
	p := person08{110,"德玛",30}
	p.PrintInfo()
	//子类可以继承父类 可以继承属性 和 方法
	//父类不能继承子类  属性  方法
	s := student08{person08{120,"德邦",15},9}

	s.PrintInfo()
}
