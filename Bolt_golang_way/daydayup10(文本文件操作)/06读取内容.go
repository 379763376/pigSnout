package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main061()  {
	//只读方式打开文件
	fp,err := os.Open("D:/a.txt")
	if err!=nil {
		fmt.Println("打开文件失败",err)
	}
	buf:= make([]byte,1024*2)   // 2k
	//块读取
	//n,_:=fp.Read(buf)
	//n 读取的长度
	//fmt.Println(string(buf[:n]))
	//n,_ = fp.Read(buf)
	//fmt.Println(string(buf[:n]))
	//n,_ = fp.Read(buf)
	//fmt.Println(string(buf[:n]))

	for {
		n,err := fp.Read(buf)
		//io.EOF文件的结尾
		//读取到文件末尾  返回值 errors.New("EOF")
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}

	//关闭文件
	defer fp.Close()
}

func main()  {
	//只读方式打开文件
	fp,err := os.Open("D:/a.txt")
	if err!=nil {
		fmt.Println("打开文件失败",err)
	}

	defer fp.Close()
	//创建文件的缓冲区
	r := bufio.NewReader(fp)
	//行读取 截取的标志 '\n'
	//slice,_ := r.ReadBytes('\n')
	//fmt.Println(string(slice))
	//slice,_ = r.ReadBytes('\n')
	//fmt.Println(string(slice))
	//slice,_ = r.ReadBytes('\n')
	//fmt.Println(string(slice))

	//for {
	//	//遇到'\n'结束读取 但是 '\n'也读取进入
	//	buf,err1 := r.ReadBytes('d')
	//	//先打印再判断
	//	fmt.Println(string(buf))
	//	if err1 != nil {
	//		if err1 == io.EOF{
	//			break
	//		}
	//		fmt.Println(err1)
	//	}
	//
	//}

	for {
		str ,err2 := r.ReadString('\n')
		fmt.Println(str)
		if err2 == io.EOF {
			break
		}
	}

}