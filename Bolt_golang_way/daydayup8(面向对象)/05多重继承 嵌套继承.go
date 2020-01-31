package main

import "fmt"

//爷
type humen struct {
	id   int
	name string
}

//父
type person05 struct {
	humen //匿名字段
	age int
	sex string
}

//子
type student05 struct {
	person05 //匿名字段
	score int
}

func main000() {
	var stu student05
	stu.name = "亚索"
	stu.sex = "男"
	stu.score = -5
	fmt.Println(stu)

	//初始化
	//自动推导类型

	stu1 := student05{person05{humen{1001,"亚索"},30,"男"},-5}
	fmt.Println(stu1)
}
type Reader interface {
	Read(p[]byte)(n int,err error)
}
type Writer interface {
	Write(p[] byte)(n int , err error)
}
type Closer interface {
	Close() error
}
type ReaderWriter interface {
	Reader
	Writer
}
type ReadCloser interface {
	Reader
	Closer
}
type WriteCloser interface {
	Writer
	Closer
}
func main() {

}