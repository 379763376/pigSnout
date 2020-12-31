package main

import (
	"fmt"
	"os"
)

func main() {
	//logger.Trace()
	//logger.Debug()
	//logger.Warning()
	//logger.Info()
	//logger.Error()

	// 写日志
	fmt.Fprintf(os.Stdout, "这是一条记录")
	fileObj,_ := os.OpenFile("./xxx.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0755)
	fmt.Fprintln(fileObj,"这是一条记录")
}