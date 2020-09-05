package main
/*
服务器和客户端通讯的5中基本类型
a.Simple strings：服务器返回的ok/pong，第一个字节是 "+", "+OK\r\n""+OK\r\n"
b.Errors：服务器返回信息，第一个字节是 "-"，"-Error message\r\n"
c.Integers：查询长度返回结果，第一个字节是 ":"，":1000\r\n"
d.Bulk Strings：命令get/lpop/hget获取到的值， 第一个字节是 "$"，"$6\r\nfoobar\r\n"
	格式：$长度\r\n具体字符\r\n
e.Arrays：发送和接收多个值，第一个字节是 "*","*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
 */
import (
	"bufio"
	"io"
)
var (
	simplePrefixSlice = []byte{'+'}
	errorsPrefixSlice = []byte{'-'}
	integerStringPrefixSlice = []byte{':'}
	bulkStringPrefixSlice = []byte{'$'}
	arrayPrefixSlice = []byte{'*'}
	lineEndingSlice = []byte{'\r','\n'}
)

//
type RESPWriter struct {
	*bufio.Writer
}

func NewRESPWriter(writer io.Writer) *RESPWriter {
	return &RESPWriter{
		Writer: bufio.NewWriter(writer),
	}
}