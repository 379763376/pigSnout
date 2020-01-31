package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var srcFileName string
	var dstFileName string
	fmt.Printf("请输入源文件名称：")
	fmt.Scan(&srcFileName)
	fmt.Printf("请输入目的文件名称：")
	fmt.Scan(&dstFileName)

	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件不能同名")
		return
	}

	//只读方式打开源文件
	sf, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("打开源文件失败", err1)
		return
	}

	//新建目的文件
	df, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("新建目的文件失败", err2)
		return
	}

	//操作完毕  关闭文件
	defer sf.Close()
	defer df.Close()

	//核心处理  从源文件读取内容  往目的文件写 读取多少 写多少
	buf := make([]byte, 1024*4) //4k

	for {
		n, err := sf.Read(buf)
		if err != nil {

			if err == io.EOF { //文件读取完毕
				break
			}
		}
		//写入
		df.Write(buf[:n])
	}

}
