package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {

}
/*
方法必须满足：
	两个参数，第二个是指针类型
	返回一个error
 */
func (p *HelloService) Hello(req string, reply *string) error {
	*reply = "hello:" + req
	return nil
}
/*
将类型实例注册为RPC服务
对象所有满足RPC规则的对象方法注册为一个RPC函数，注册在HS服务空间下
建立一个TCP连接
通过rpc.ServerConn在TCP连接上提供RPC服务
 */
func main() {
	rpc.RegisterName("HS", new(HelloService))
	listener,err := net.Listen("tcp", ":1234")
	if err != nil{
		log.Fatal("ListenTCP error:", err)
	}
	conn,err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
