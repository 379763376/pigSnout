package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)
//1. 一个包里不能有重复名称的函数
//func run()  {
//	fmt.Println("run")
//}

func func1(args ...int)  {
	//不能将不定参的名称传给另一个不定参，如果要传值需要获取指定个数的数据
	//func2(args) 错误写法
	func2(args[0],args[1],args[2])
	func2(args[0:len(args)]...)//[:3] [2:3]
	func2(args[:]...)
	func2(args...)
}
func func2(args ...int)  {
	fmt.Println(args)
}
func main() {
	//2.不定长参数的调用
	func1(1,2,3,4,5,6)
	//3.函数的嵌套使用 在一个函数中调用另一个函数
	//4.函数嵌套的执行过程
	t1(10,20)
	//5. 函数返回值
	var result = func3()
	var result2  =  func4(1,5)
	fmt.Println(result,result2)
	//6. 函数返回多个值
	a,b,c := func5(1,5)
	fmt.Println(a,b,c)
	//6.匿名变量  丢弃接收的数据
	_,_,sum := func5(1,5)
	fmt.Println(sum)
}
func t1(x1,x2 int){
	t2(x1,x2)
}
func t2(x1,x2 int)  {
	sum := x1+x2
	fmt.Println(sum)
}
func func3() int {  //返回整型数据  通过return返回数据
	var a int = 5
	var b int =1
	var sum = a+ b
	return sum
}
func func4(a,b int) (sum int) {  //返回整型数据  通过return返回数据
	sum = a+ b
	return //sum
}
func func5(a,b int) (a1,b1,sum int) {  //返回多个值
	sum = a+ b
	a1=a
	b1=b
	return //sum
}
func func55(a,b int) (int, int) {  //返回值的一种写法
	sum2 := a+ b
	return sum2,a
}


func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) { //
	return a / b, a % b
}
//函数作为参数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer() //获得函数的指针
	opName := runtime.FuncForPC(p).Name() //获得函数名字
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main02() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))//计算3的4次方

	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}