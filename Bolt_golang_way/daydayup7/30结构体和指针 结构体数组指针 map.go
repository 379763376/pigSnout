package main

import "fmt"

type person struct {
	id   int
	name string
	age  int
}

func main2401() {
	//初始化
	var per person =person{101,"李宁",40}
	//fmt.Println(per)
	//fmt.Printf("%p\n",&per)

	//定义指针接收结构体变量地址
	//p := &per
	var p *person = &per
	fmt.Printf("%T\n",p)    //*person 类型

	//通过指针间接修改结构体成员的值
	(*p).age = 50
	fmt.Println(per)

	//指针直接操作结构体成员
	p.id = 110
	fmt.Println(per)

}
//结构体指针作为函数参数
func test24(p *person)  {
	p.age = 57
}


func main2402()  {

	var per person = person{101,"李宁",40}
	//地址传递
	test24(&per)
	fmt.Println(per)
}

func main2403()  {
	arr := [3]person{{101,"钢铁侠",34},
		{102,"绿巨人",40},
		{103,"黑寡妇",28}}
	//指向结构体数组的指针
	p := &arr
	//fmt.Printf("%p\n",p)
	//fmt.Printf("%T\n",p)

	p[0].age = 40//ok
	(*p)[0].age = 39

	fmt.Println(arr[0])

	for i:=0;i<len(p) ;i++  {
		fmt.Println(p[i])
	}


}

func main()  {
	//map类型变量
	m := make(map[int]*[3]person)

	fmt.Printf("%T\n",m)

	m[1] = new([3]person)
	m[1] = &[3]person{{101,"钢铁侠",34},
		{102,"绿巨人",40},
		{103,"黑寡妇",28}}

	m[2] = new([3]person)
	m[2] = &[3]person{{101,"美队",34},
		{102,"黑豹",40},
		{103,"女巫",28}}


	for k,v := range m{
		fmt.Println(k,*v)
	}


	//数组指针
	var p *[3]int

	//创建内存空间 存储 [3]int
	p = new([3]int)
	p[0] = 123
	p[1] = 456
	p[2] = 789

	fmt.Println(p)


}