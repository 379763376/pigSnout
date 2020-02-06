package main

import "fmt"

type hero struct {
	name  string
	age   int
	power int
}

//函数参数
func test17(m map[int]hero) {

	//err
	//m[102].power = 89

	stu := m[102]
	stu.power = 89
	m[102] = stu

	fmt.Println(m)
	fmt.Printf("%p\n", m)
}

func main1701() {
	//将结构体作为map中的值  value
	m := make(map[int]hero)

	//map中的数据不建议排序操作
	m[101] = hero{"钢铁侠", 30, 100}
	m[102] = hero{"美队", 30, 90}

	fmt.Println(m)

	delete(m, 102)
	fmt.Println(m)

}

func main1702() {
	//value 类型 是一个切片
	m := make(map[int][]hero)

	m[101] = []hero{{"钢铁侠", 30, 100},
		{"蜘蛛侠", 17, 80}}

	m[101] = append(m[101], hero{"星爵", 30, 10})
	fmt.Println(m)

	m[102] = []hero{{"美队", 30, 90}}
	m[102] = append(m[102], hero{"冬兵", 30, 75})
	fmt.Println(m[102][0])
}

func main() {
	m := make(map[int]hero)
	//map中的数据不建议排序操作
	m[101] = hero{"钢铁侠", 30, 100}
	m[102] = hero{"美队", 30, 90}
	//将map作为函数参数   地址传递
	test17(m)

	fmt.Println(m)

	fmt.Printf("%p\n", m)
}
