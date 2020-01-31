package main

import "fmt"

/*
反射可以提高程序的灵活性 使得interface{}有更大的发挥余地
反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
反射会将你名字段作为独立字段（匿名字段本质）
如果要用反射修改对象状态 interface.data需要是settable 即 pointer-interface
射可以动态调用方法反
*/

//========测试传入数据类型========
func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	default:
		return "???"
	}
}

func main01() {
	fmt.Println(Sprint("输入的数据类型可能很复杂，switch无法预测"))
}
//=====================================
/*

 */