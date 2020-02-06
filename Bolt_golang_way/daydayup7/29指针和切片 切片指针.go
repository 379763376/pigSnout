package main

import "fmt"

func main2301() {

	s := []int{1,2,3,4,5}
	//指针和切片建立联系
	p := &s
	// b := *p == s
	fmt.Printf("%p\n",p) //存放切片名的地址
	fmt.Printf("%p\n",s) //切片存放在一个内存空间 切片名存放切片的地址
	//*[]int 切片指针
	//var p *[]int  指针切片
	fmt.Printf("%T\n",p)


	//通过指针间接操作切片元素
	//p[1] = 222 通过p找到的是s的存放地址
	(*p)[1] = 222  //p代表s的地址 s的地址存放的是切片的地址 *p指向切片

	fmt.Println(s)
	//for循环遍历
	for i:=0;i<len(s) ; i++ {
		fmt.Println((*p)[i])
	}

}
//切片指针作为函数参数
func test23(s *[]int)  {  //形参s存放的是实参s在栈区的地址
	*s = append(*s,6,7,8,9)
}


func main()  {
	s := []int{1,2,3,4,5}  //切片名本身就是一个地址  切片存放在堆区 切片名存放在栈区
	//地址传递
	test23(&s)
	fmt.Println(s)
}