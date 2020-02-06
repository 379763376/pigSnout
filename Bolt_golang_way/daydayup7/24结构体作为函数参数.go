package main

import "fmt"

type Hero struct {
	name  string
	age   int
	power int
}
//结构体作为函数参数 是值传递
func test18(h Hero)  {
	h.power = 120
	fmt.Println(h)
}


func main() {
	//结构体变量
	h := Hero{"钢铁侠",30,100}

	//值传递
	test18(h)

	fmt.Println(h)
}
