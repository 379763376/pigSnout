package main
<<<<<<< HEAD
=======

>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
import (
	"fmt"
)

type A struct {
	Name string
<<<<<<< HEAD
	age int
}
type B struct {
	Name string
=======
	age  int
}
type B struct {
	Name  string
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	Score float64
}
type C struct {
	A
	B
	//Name string
}

type D struct {
	a A //有名结构体
}

<<<<<<< HEAD

type Goods struct {
	Name string
=======
type Goods struct {
	Name  string
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	Price float64
}

type Brand struct {
<<<<<<< HEAD
	Name string
=======
	Name    string
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	Address string
}

type TV struct {
	Goods
<<<<<<< HEAD
	Brand	
=======
	Brand
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
}

type TV2 struct {
	*Goods
<<<<<<< HEAD
	*Brand	
}

type Monster struct  {
	Name string
	Age int
=======
	*Brand
}

type Monster struct {
	Name string
	Age  int
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
}

type E struct {
	Monster
	int
	n int
}

func main() {
	var c C
	//如果c 没有Name字段，而A 和 B有Name, 这时就必须通过指定匿名结构体名字来区分
	//所以 c.Name 就会包编译错误， 这个规则对方法也是一样的！
	c.A.Name = "tom" // error
	fmt.Println("c")

	//如果D 中是一个有名结构体，则访问有名结构体的字段时，就必须带上有名结构体的名字
<<<<<<< HEAD
	//比如 d.a.Name 
	var d D 
	d.a.Name = "jack"


	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{ Goods{"电视机001", 5000.99},  Brand{"海尔", "山东"}, }

	//演示访问Goods的Name
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price) 

	tv2 := TV{ 
		Goods{
			Price : 5000.99,
			Name : "电视机002", 
		},  
		Brand{
			Name : "夏普", 
			Address :"北京",
		}, 
=======
	//比如 d.a.Name
	var d D
	d.a.Name = "jack"

	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}}

	//演示访问Goods的Name
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)

	tv2 := TV{
		Goods{
			Price: 5000.99,
			Name:  "电视机002",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	}

	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)

<<<<<<< HEAD
	tv3 := TV2{ &Goods{"电视机003", 7000.99},  &Brand{"创维", "河南"}, }

	tv4 := TV2{ 
			&Goods{
				Name : "电视机004", 
				Price : 9000.99,
			},  
			&Brand{
				Name : "长虹", 
				Address : "四川",
			}, 
		}
=======
	tv3 := TV2{&Goods{"电视机003", 7000.99}, &Brand{"创维", "河南"}}

	tv4 := TV2{
		&Goods{
			Name:  "电视机004",
			Price: 9000.99,
		},
		&Brand{
			Name:    "长虹",
			Address: "四川",
		},
	}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab

	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)

<<<<<<< HEAD

=======
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	//演示一下匿名字段时基本数据类型的使用
	var e E
	e.Name = "狐狸精"
	e.Age = 300
	e.int = 20
	e.n = 40
	fmt.Println("e=", e)

<<<<<<< HEAD
}
=======
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
