package main

import (
	"fmt"
	"os"
)

func main() {

	//os.Open  只读方式打开
	//fp,err := os.Open("D:/a.txt")

	//os.OpenFile(文件名，打开方式，打开权限)
	fp,err := os.OpenFile("D:/a.txt",os.O_RDWR,6)
	if err!=nil {
		fmt.Println("打开文件失败")
	}

	fp.WriteString("hello")
	fp.WriteAt([]byte("hello"),25)


	defer fp.Close()
}
