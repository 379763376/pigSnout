package main
<<<<<<< HEAD
import (
	"fmt"
)
=======

import (
	"fmt"
)

>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
type Point struct {
	x int
	y int
}
<<<<<<< HEAD
func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point  //oK
=======

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //oK
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	// 如何将 a 赋给一个Point变量?
	var b Point
	// b = a 不可以
	// b = a.(Point) // 可以
<<<<<<< HEAD
	b = a.(Point) 
	fmt.Println(b) // 

=======
	b = a.(Point)
	fmt.Println(b) //
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab

	//类型断言的其它案例
	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2  //空接口，可以接收任意类型
	// // x=>float32 [使用类型断言]
	// y := x.(float32)
	// fmt.Printf("y 的类型是 %T 值是=%v", y, y)

<<<<<<< HEAD

	//类型断言(带检测的)
	var x interface{}
	var b2 float32 = 2.1
	x = b2  //空接口，可以接收任意类型
=======
	//类型断言(带检测的)
	var x interface{}
	var b2 float32 = 2.1
	x = b2 //空接口，可以接收任意类型
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	// x=>float32 [使用类型断言]

	//类型断言(带检测的)
	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")

}
