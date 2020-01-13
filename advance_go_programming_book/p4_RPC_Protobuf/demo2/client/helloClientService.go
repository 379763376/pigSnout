package main

import (
	"net/rpc"
	"pigSnout/advance_go_programming_book/p4_RPC_Protobuf/demo2/common"
)

/*
简化客户端用户调用RPC函数，接口规范部分增加对客户端的简单包装
*/
type HelloServiceClient struct {
	*rpc.Client
}

//var _HelloServiceInterface = (*HelloServiceClient)(nil)
func DialHelloService(network, address string) (*HelloServiceClient, error){
	c,err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(req string, reply *string) error  {
	return p.Client.Call(common.HelloServiceName+".Hello", req, reply)
}