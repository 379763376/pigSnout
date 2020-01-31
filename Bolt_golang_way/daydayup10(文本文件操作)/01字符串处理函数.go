package main
/*
strings包
 */
import (
	"fmt"
	"strings"
)

func main0101() {
	//查找一个字符串在另一个字符串中是否出现
	str1 := "hello world"
	str2 := "g"

	//Contains(被查找的字符串，查找的字符串)  返回值 bool
	//一般用于模糊查找
	b := strings.Contains(str1,str2)
	//fmt.Println(b)
	if b {
		fmt.Println("找到了")
	}else {
		fmt.Println("没有找到")
	}

}

func main0102()  {
	//字符串切片
	slice := []string{"123","456","789"}
	//fmt.Println(slice)
	//Join
	//字符串的连接
	str := strings.Join(slice,"")
	fmt.Println(str)
	//fmt.Printf("%T\n",str)
}

func main0103()  {
	str1 := "hello world"
	str2 := "e"
	//查找一个字符串在另一个字符串中第一次出现的位置 返回值  int  下标  -1 找不到
	
	i := strings.Index(str1,str2)
	fmt.Println(i)
}

func main0104()  {
	str := "性感网友，在线取名。"
	//将一个字符串重复n次
	str1 := strings.Repeat(str,100)
	fmt.Println(str1)

}

func main0105()  {
	str := "性感网友在线取名性感性感性感性感性感"
	//字符串替换  屏蔽敏感词汇
	//如果替换次数小于0 表示全部替换
	str1 := strings.Replace(str,"性感","**",-1)
	fmt.Println(str1)

}

func main0106()  {
	//str1 := "1300-188-1999"
	//将一个字符串按照标志位进行切割变成切片
	str1 := "123456789@qq.com"
	slice := strings.Split(str1,"@")
	fmt.Println(slice[0])
}

func main0107()  {

	str := "====a===u=ok===="
	//去掉字符串头尾的内容
	str1:= strings.Trim(str,"=")
	fmt.Println(str1)
	
}

func main()  {
	str := "    are you ok    "
	//去除字符串中空格  转成切片  一般用于统计单词个数
	slice := strings.Fields(str)
	fmt.Println(slice)
}

//总结

	//查找
	//1.bool类型 := strings.Contains(被查找字符串，查找字符串)
	//2.int类型 := strings.Index(被查找字符串，查找字符串)

	//分割
	//[]string类型 := strings.Spilt(切割字符串，标志)

	//组合
	//string类型 := strings.Join(字符串切片，标志)

	//重复
	//string类型 := strings.Repeat(字符串，次数)

	//替换
	//string类型 := strings.Replace(字符串，被替换字符串，替换字符串，次数)

	//去掉内容
	//string类型 := strings.Trim(字符串，去掉字符串)
	//[]string类型 := strings.Fields(字符串)