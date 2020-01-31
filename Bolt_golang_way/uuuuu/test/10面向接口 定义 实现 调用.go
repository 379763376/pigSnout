package main

import (
	"Bolt_golang_way/uuuuu/test/mock"
	"Bolt_golang_way/uuuuu/test/real"
	"fmt"
	"time"
)
/*
1.使用者定义接口
接口
 */
type Retriver interface {
	Get(url string) string
}//接口Get
type Poster interface {
	Post(url string,
		form map[string]string) string
}//接口Post
/*
2.使用者使用接口
接口方法
 */
func download(r Retriver) string{
	return r.Get("http://www.baidu.com")
} //使用download
func post(poster Poster)  string {
	return poster.Post("http://www.baidu.com",
		map[string]string{
			"name":"bolt",
			"age":"23",
		})
}//使用post

//接口一 接口二 组合
type RetriverPoster interface {
	Retriver
	Poster
}
func session(s RetriverPoster) string {
	//s.Get()
	s.Post("http://www.baidu.com", map[string]string{
		"contents":"接口调用时修改初始实现时的值",
	})
	return s.Get("http://www.baidu.com")
}
func main() {


	//3.定义结构对象，并实例化，调用接口方法
	var r Retriver
	r = mock.Retriver1{"this is my interface test1"} //实现者对象（实现接口 只要实现了Get方法)
	fmt.Println(download(r))//调用接口方法（传递实现者对象）
	fmt.Println(download(mock.Retriver1{"this is two"}))
	r = &mock.Retriver1{"this is my interface test1"}//自带指针 对应的接收者可以指针可以值
	r = &real.Retriver2{} //实现者实现了接口,是一个指针接收者
	//fmt.Println(download(r)) //使用者使用
	r = &real.Retriver2{
		"Mozilla/5。0",
		time.Minute,
	}
	/*
	4.接口对象中包含了类型和值
	判断类型的两种方式
	 */
	inspect(r)//判断类型，方法一：switch case
	//判断类型，方法二：type assertion r.(类型)取得类型
	realRetriver := r.(*real.Retriver2)
	fmt.Println(realRetriver.UserAgent)
	//panic: interface conversion: main.Retriver is *real.Retriver2, not mock.Retriver1
	//r.(类型)  是把接口变量肚子里的值强制类型转换，转换成功就是ok,否则err
	if mockRetriver,ok := r.(mock.Retriver1);ok{
		fmt.Println(mockRetriver.Contents)
	}else{
		fmt.Println("not a mock retriver")
	}
	/*
	5.改造queue接收任何类型
	type Queue []interface{}
	设置只接收int类型
	*q = append(*q,v.(int)) v是接口类型变量 包含类型和数据两部分 强制转为int
	 */
	/*
	6.接口方法组合调用
	 */
	//不能直接使用r 因为r只有接口的能力，只有get,没有post
	retriver := mock.Retriver1{
		"接口实现中定义的content"} //结构体创建了 具有了get post方法
	fmt.Println("try a session")
	//fmt.Println(session(retriver))
	//执行结构体对象的post(是值传递，方法内部修改了拷贝的结构体，不会改变结构体对象本身)
	//执行get打印输出
	fmt.Println(session(&retriver))

	/*
	7.
	 */
}

func inspect(r Retriver) {
	fmt.Println("Inspecting",r)
	fmt.Printf(" > %T %v\n",r,r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case mock.Retriver1:
		fmt.Println("contents:", v.Contents)
	case *real.Retriver2:
		fmt.Println("UserAgent", v.UserAgent)
	}
	fmt.Println()

}
