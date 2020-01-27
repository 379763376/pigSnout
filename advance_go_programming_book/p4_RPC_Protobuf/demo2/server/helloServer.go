package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

/*
方法必须满足：
	两个参数，第二个是指针类型
	返回一个error
*/
func (p *HelloService) Hello(req string, reply *string) error {
	*reply = "hello:" + req
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}

}
