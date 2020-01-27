package main

import (
	"net/rpc"
	"pigSnout/advance_go_programming_book_bolt/p4_RPC_Protobuf/demo2/common"
)

/*
明确服务名字和接口
*/

type HelloServiceInterface = interface {
	Hello(req string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(common.HelloServiceName, svc)
}


