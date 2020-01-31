package main
/*
Go 编译型语言
go语言工具将代码及依赖转换成计算机的机器指令（静态编译）
run 可以编译一个或多个以go结尾的源文件，链接库文件，生成可执行文件
$ go run helloworld.go
go语言支持原生Unicode，处理全世界任何语言的文本
$ go build helloworld.go 编译保存，windows生成exe文件，可以直接执行
$ ./helloworld
=================================
package
main 特殊的包 定义了一个独立可执行的程序，而不是一个库
import 只能引入需要的包
gofmt 工具 格式化代码
goimports  工具  根据需求自动导包，删除import
===============================
在每个包的声明前添加注释，整体描述
=============================
os.Args提供外部值 是一个切片
os.Args[0]是命名本身，os.Args[1:]实际传递的值
===============================
for 循环中的三个部分 每个都可以省略
for range
================================
string包Join函数
*/
import (
	"fmt"
	"os"

	//"bufio"
	//"io"
)

func myf2()  {
	
}

func main01() {
	fmt.Printf("hello, 世界\n")
	//fmt.Printf(stringutil.Reverse("!oG,olleH"))
	//reader := bufio.Reader{}
	//fmt.Println(reader)
	//reader2 := io.Reader([byte(1)])
	//fmt.Println(reader2)
}
func main() {
	var s,sep string
	for i :=1;i<len(os.Args);i++  {
		s+=sep + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}