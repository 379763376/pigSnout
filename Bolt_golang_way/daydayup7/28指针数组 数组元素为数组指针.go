package main

import "fmt"

func main() {

	a := 10
	b := 20
	c := 30

	//指针数组  数组元素 是指针类型
	var arr [3]*int = [3]*int{&a, &b, &c}


	//数组指针 存放的是数组的地址  *[3]int   [3]int的指针
	//指针数组 存放的是一些变量的地址  [3]*int  指针类型的数组


	fmt.Println(arr)

	*arr[1] = 200 //先arr[1]值是一个地址，*arr[1]代表地址arr[1]的值 赋值为200

 	fmt.Println(b)

	for i := 0; i < len(arr); i++ {
		fmt.Println(*arr[i])
	}
}

func main01() {
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	c := [3]int{7, 8, 9}

	//p := &a
	//fmt.Println(p)
	//数组中的元素是  数组指针    定义了数组[3]类型为数组指针*[3]int
	var arr [3]*[3]int = [3]*[3]int{&a,&b,&c}

	fmt.Println(arr)
	fmt.Printf("%T\n",arr)

	(*arr[1])[1] = 555  //*arr[1] 为
	arr[1][2] = 666
	fmt.Println(b)


}
