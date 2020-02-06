package main

import (
	"log"
	"net"
	"net/rpc"

	pb "chai2010.cn/gobook/examples/ch4.2/hello.pb"
)

type HelloService struct{}

func (p *HelloService) Hello(request pb.String, reply *pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func main() {
	rpc.Register(new(HelloService))

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
