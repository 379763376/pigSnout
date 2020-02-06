package main
/*
go 不允许隐式转换 都需要声明
strconw字符串与其他类型转换的包
 */
import (
	"fmt"
	"strconv"
)

func main0201() {

	str := "hello world"
	//将字符串转成字符切片   强制类型转换
	slice := []byte(str)

	fmt.Println(slice)
}

func main0202()  {
	//字符切片

	slice := []byte{'h','e','l','l','o',97}
	fmt.Println(slice)

	fmt.Println(string(slice))
}

func main0203()  {
	//将其他类型转成字符串  Format
	//b := false
	//str := strconv.FormatBool(true)
	//fmt.Println(str)
	//fmt.Printf("%T\n",str)

	//str := strconv.FormatInt(120,10)  //计算机中进制 可以表示2-36  2  8  10 16
	//fmt.Println(str)

	//str:= strconv.FormatFloat(3.14159,'f',4,64)
	//fmt.Println(str)

	str := strconv.Itoa(123)
	fmt.Println(str)
}

func main0204()  {
	//字符串转成其他类型  Parse

	//b,err := strconv.ParseBool("true")
	//if err!=nil {
	//	fmt.Println("类型转换出错")
	//
	//}else {
	//	fmt.Println(b)
	//	fmt.Printf("%T\n",b)
	//}
	//字符串转换成16进制整型
	//v,err := strconv.ParseInt("abc",16,64)
	//fmt.Println(v,err)

	//v,_ := strconv.ParseFloat("3.14159",64)
	//fmt.Println(v)


	//等价于ParseInt(s, 10, 0)
	v,_:=strconv.Atoi("123")
	fmt.Println(v)

}

func main()  {

	slice := make([]byte,0,1024)
	//将其他类型转成字符串 添加到字符切片里

	slice = strconv.AppendBool(slice,false)
	slice = strconv.AppendInt(slice,123,2)
	slice = strconv.AppendFloat(slice,3.14159,'f',4,64)
	slice = strconv.AppendQuote(slice,"hello")
	fmt.Println(string(slice))
}