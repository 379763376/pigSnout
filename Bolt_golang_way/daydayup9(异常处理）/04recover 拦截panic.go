package main
/*
从panic恢复 获得程序的控制权
 */
import "fmt"

func testA()  {
	fmt.Println("testA")

}

func testB(x int)  {
	//设置recover()

	//在defer调用的函数中使用recover()
	defer func() {
		//防止程序崩溃
		//recover()
		//fmt.Println(recover())

		if err:=recover();err!=nil {
			fmt.Println(err)
		}
	}()  //匿名函数

	var a [3]int
	a[x] = 999
}

func testC()  {
	fmt.Println("testC")
}
func main() {
	testA()
	testB(0)  //发生异常 中断程序
	testC()
}
