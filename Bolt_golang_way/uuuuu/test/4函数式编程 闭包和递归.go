 package main

import "fmt"

func main() {
	sum := 0
	//1.匿名函数 这个函数使用的sum是返回值定义的变量 而不是外部函数变量
	x,y,z := func(a,b int)(i,j,sum int){
		i= a
		j =b
		sum = a+b
		return
	}(1,3)
	fmt.Println(x,y,z,sum)
	//2.匿名函数 这个函数使用的sum是返回值定义的变量 而不是外部函数变量
	func(a,b int){
		sum = a+b
		return
	}(1,3)
	fmt.Println(sum)
	//3匿名函数定义方式
	//3.1
	f := func() int{
		sum ++
		return 1
	}
	f()
	//3.2
	fmt.Println(func(){
		sum++
	})
	//3.3 把一个匿名函数赋值给另一个
	type anonymityFuncType func() int
	var anonymityFunc anonymityFuncType
	anonymityFunc = f
	a := anonymityFunc()
	fmt.Println(a)

	//4.递归
	b := add100(100)
	fmt.Println(b)

	//5.函数式编程 从0加到10
	af:= adder()
	for i :=0 ;i<10;i++{
		fmt.Printf("$%d,",af(i))
	}
	fmt.Println()
	//6.正统 函数式编程  只有常量和函数 不能有if和循环==》所以状态放到新的函数
	//函数 返回当前值和下一轮要执行的函数
	al := adder2(0)
	for i:=0;i<10 ;i++  {
		sum,al = al(i)
		fmt.Printf("%d,",sum)
	}
}
func add100(num int)  int{
	if num == 1{
		return 1
	}
	return num+add100(num-1)
}
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//函数类型
type iAdder func(int)(int,iAdder)
//函数
 func adder2(base int) iAdder  {
	 return func(v int) (int,iAdder) {
		 return base+v,adder2(base+v)
	 }
 }