package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)
//1.不定长参数
func test1(t...int) int{
	return  1
}
//不定参要放最后
func test2(x string,y...int){
//给不定参的名称传递给另一个不定参
	test1(y...)
	test1(y[:]...)
	test1(y[2:3]...)
	test1(y[0:len(y)]...)
	test1(y[2],y[3])
}

//2.函数作为参数 函数参数在其他类型参数前面
func apply(op func(int,int) int,a,b int) int{
	p := reflect.ValueOf(op).Pointer()//得到函数的指针
	opName := runtime.FuncForPC(p).Name() //获得函数的名称
	fmt.Printf("calling function %s with args" +
		"(%d,%d)\n",opName,a,b)
	return op(a,b)
}
func main() {
	test1(1,2,3,4,5)
	test2("a",1,2,3,4,5)

	//传入三个参数：计算函数 数值1 数值2
	fmt.Println("pow (3,4) is :",apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a),float64(b)))
		},3,4))

	var b myFuncType
	b = test1 //两者的参数和返回值类型一致 就可以赋值
	b(0,1,2,3)
}
//函数类型 和 给已存在的函数起别名方法一直
type myFuncType func(x...int) int
type bigint int64
