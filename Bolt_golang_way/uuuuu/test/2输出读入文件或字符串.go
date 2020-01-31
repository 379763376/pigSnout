package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//读取文件输出
	printFile("abc.txt")
	//读取字符串 转换为Reader
	s := `abc"d"
	kkkk
	123

	p`
	printFileContents(strings.NewReader(s))


}
func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	printFileContents(file)

}
