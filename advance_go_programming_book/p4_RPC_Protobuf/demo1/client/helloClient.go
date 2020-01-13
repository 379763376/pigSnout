package main

import (
	"fmt"
	"log"
	"net/rpc"
)
/*
rpc.Dial拨号RPCRPC服务
client.Call第一个参数是服务命名空间和方法名，第二、三是方法参数
 */
func main() {
	client,err := rpc.Dial("tcp", "localhost:1234")
	if err != nil{
		log.Fatal("dialing:",err)
	}
	var reply string
	err = client.Call("HS.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply)
}
