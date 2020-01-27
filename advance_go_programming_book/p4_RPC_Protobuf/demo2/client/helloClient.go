package main

import (
	"log"
)
/*
调用时用接口中的服务名字
 */
func main() {
	client,err := DialHelloService("tcp", "localhost:1234")
	if err != nil{
		log.Fatal("dialing:",err)
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	println(reply)
}
