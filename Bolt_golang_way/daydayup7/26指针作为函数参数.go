package main

import "fmt"
//指针作为函数参数
func Swap(num1,num2 *int)  {
	*num1,*num2 = *num2,*num1
	fmt.Println(*num1,*num2)
}

func main() {
	a := 10
	b := 20

	//&变量  取地址操作 引用运算符
	//*指针变量   取值操作  解引用运算符
	Swap(&a,&b) //地址传递
	fmt.Println(a,b)



}
