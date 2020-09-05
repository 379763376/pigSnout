package main
//
//import (
//	"go.etcd.io/etcd/clientv3"
//)
//
//config := clientv3.Config{
//	Endpoints:[]string{"10.164.25.205:30379","10.164.25.223:30379"},
//	DialTimeout:10*time.Second,
//}
//client,err := clientv3.New(config)
//if err != nil {
//	panic(err)
//}
//defer client.Close()
//kv := clientv3.NewKV(client)