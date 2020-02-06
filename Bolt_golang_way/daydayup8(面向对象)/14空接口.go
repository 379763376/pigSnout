package main

import "fmt"

func main1401() {

	var i interface{}
	//接口类型可以接收任意类型的数据
	//fmt.Println(i)
	fmt.Printf("%T\n",i)
	i = 10
	fmt.Println(i)
	fmt.Printf("%T\n",i)

	i = 3.14
	fmt.Println(i)
	fmt.Printf("%T\n",i)

	i = "Go"
	fmt.Println(i)
	fmt.Printf("%T\n",i)

	fmt.Println("aaa",11,3.14,'a')

}

func test14(){
	fmt.Println("test14")
}

func main()  {
	//空接口类型的切片
	var i []interface{}
	fmt.Printf("%T\n",i)
	i = append(i,10,3.14,"aaa",test14)

	fmt.Println(i)
	for idx:=0; idx<len(i) ;idx++  {
		fmt.Println(i[idx])

	}
}