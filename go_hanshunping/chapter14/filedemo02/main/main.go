package main
<<<<<<< HEAD
import (
	"fmt"
	"os"
	"bufio"
	"io" 
)
=======

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
func main() {
	//打开文件
	//概念说明: file 的叫法
	//1. file 叫 file对象
	//2. file 叫 file指针
	//3. file 叫 file 文件句柄
<<<<<<< HEAD
	file , err := os.Open("d:/test.txt")
=======
	file, err := os.Open("d:/test.txt")
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏.

	// 创建一个 *Reader  ，是带缓冲的
	/*
<<<<<<< HEAD
	const (
	defaultBufSize = 4096 //默认的缓冲区为4096
	)
=======
		const (
		defaultBufSize = 4096 //默认的缓冲区为4096
		)
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	*/
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
<<<<<<< HEAD
		if err == io.EOF { // io.EOF表示文件的末尾
=======
		if err == io.EOF {                  // io.EOF表示文件的末尾
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
			break
		}
		//输出内容
		fmt.Printf(str)
	}

	fmt.Println("文件读取结束...")
<<<<<<< HEAD
}
=======
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
