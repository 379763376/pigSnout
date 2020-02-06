package main

import "fmt"

func main2501() {
	a := 10
	b := 20
	// 一级指针  指向变量的地址
	p := &a
	fmt.Println(*p)
	//fmt.Printf("%T\n", p)
	//二级指针  指向一级指针的地址
	var pp **int
	//var pp ***int
	//var pp ****int
	pp = &p
	//通过二级指针链接修改一级指针的值
	//*pp = &b
	//通过二级指针间接修改变量的值
	**pp = 100
	//fmt.Printf("%T\n", pp)
	fmt.Println(*p)
	fmt.Println(a)
	fmt.Println(b)
}

func main()  {
	a := 10
	var p *int = &a
	var pp **int = &p
	var ppp ***int = &pp

	fmt.Println(ppp)

	//引用运算符 不能连续使用

	//***ppp = **pp = *p =a
}