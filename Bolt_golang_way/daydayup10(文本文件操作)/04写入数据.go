package main

import (
	"fmt"
	"io"
	"os"
)

func main0401() {
	//os.Create(文件名) 文件名  可以写绝对路径和相对路径
	//返回值  文件指针 错误信息

	//\反斜杠 转义字符
	//在写路径时可以使用/正斜杠代替\反斜杠
	fp,err := os.Create("D:/a.txt")
	if err!=nil{
		//文件创建失败
		/*
		1.路径不存在
		2.文件权限
		3.程序打开文件上限
		 */
		fmt.Println("文件创建失败")
		return
	}

	//写文件

	//\n不会换行  原因 在windows文本文件中换行\r\n  回车  在linux中换行\n
	fp.WriteString("hello world\r\n")
	fp.WriteString("性感荷官在线发牌")


	defer fp.Close()

	//关闭文件
	//如果打开文件不关闭 造成内存的浪费  程序打开文件的上限
	//fp.Close()
}

func main0402()  {
	fp,err := os.Create("D:/a.txt")
	if err!=nil{
		//文件创建失败
		/*
		1.路径不存在
		2.文件权限
		3.程序打开文件上限
		 */
		fmt.Println("文件创建失败")
		return
	}


	//写操作

	//slice := []byte{'h','e','l','l','o'}


	//count,err1 := fp.Write(slice)
	count,err1 := fp.Write([]byte("性感老王在线授课"))


	if err1!=nil {
		fmt.Println("写入文件失败")
		return
	}else {
		fmt.Println(count)
	}

	defer fp.Close()
}

func main()  {
	fp,err := os.Create("D:/a.txt")
	if err!=nil{
		//文件创建失败
		/*
		1.路径不存在
		2.文件权限
		3.程序打开文件上限
		 */
		fmt.Println("文件创建失败")
		return
	}

	//写操作
	//获取光标流位置'
	//获取文件起始到结尾有多少个字符
	//count,_:=fp.Seek(0,os.SEEK_END)
	count,_:=fp.Seek(0,io.SeekEnd)
	fmt.Println(count)
	//指定位置写入
	fp.WriteAt([]byte("hello world"),count)
	fp.WriteAt([]byte("hahaha"),0)
	fp.WriteAt([]byte("秀儿"),19)

	defer fp.Close()
}