package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//已有的算法
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//定义一个函数类型类型 和斐波那契数列函数返回的一致
type intGen func() int
//目的 通过Reader使用斐波那契数列函数
func (fib intGen) Read(p []byte) (n int, err error) {//n是read读了几个字节 error出了哪些错
	next := fib()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next) //写Read太底层，找实现了Read的Sprintf帮我们读一下

	// TODO: incorrect if p is too small!
	//数据太小不给读
	return strings.NewReader(s).Read(p)//调用接口方法 传递一个实现了方法的对象
}

//函数类型实现Reader接口的Read方法
func main() {
	var fib intGen= Fibonacci()
	pFileContents(fib)//intGen实现了Read方法 去不断的读入数据 直到文件接收io.EOF
}
func pFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}