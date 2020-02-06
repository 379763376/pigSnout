package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Create(文件名) 文件名  可以写绝对路径和相对路径
	//返回值  文件指针 错误信息
	fp,err := os.Create("./a.txt")
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

	//读写文件

	defer fp.Close()

	//关闭文件
	//如果打开文件不关闭 造成内存的浪费  程序打开文件的上限
	//fp.Close()
}
