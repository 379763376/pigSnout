package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件
//Open 只能读
//Openfile 以指定模式打开文件，从而实现写入相关功能

func main() {
	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
//file.Read
func readFromFile()  {
	fileObj,err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//读文件
	//var tmp = make([]byte,128) //指定读的长度
	var tmp  [128]byte
	for {
		n,err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v",err)
			return
		}
		fmt.Println("读了%d个字节\n",n)
		fmt.Println(string(tmp[:n]))
		if n < 0 {
			return
		}
	}
}



// 使用bufio.NewReader
func readFromFileByBufio()  {
	fileObj,err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line,err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v",err)
		}
		fmt.Print(line)
	}

}

//ioutil读取整个文件
func readFromFileByIoutil()  {
	ret,err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}


// 写Write和WriteString
//os.O_TRUNC 清空重新写
func writeByWriter()  {
	fileobj,err :=os.OpenFile("./xx.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	if err != nil{
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	fileobj.Write([]byte("XXxx\n"))
	fileobj.WriteString("我是中文")
	fileobj.Close()
}
//bufio.NewWriter
func writeByBufio()  {
	fileobj,err :=os.OpenFile("./xx.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	if err != nil{
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileobj.Close()
	writer := bufio.NewWriter(fileobj)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}
//ioutil.WriteFile直接往文件写东西
func writeByIotil()  {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}


// CopyFile 拷贝文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') //注意是字符
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

//从命令行读入输出
// fmt.Scanln(&s) 读取的数据中间不能有空格  包含空白符：空白符包含空格和换行
func useBufio()  {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s,_ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n",s)
}