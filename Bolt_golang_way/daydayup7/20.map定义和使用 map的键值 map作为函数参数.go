/*
map datatype
字典结构 map
 */
package main

import "fmt"

func main() {
	//一。定义
	//1.定义方式1
	//map 【键类型】 值类型
	var m1 map[int]string
	fmt.Println(m1)
	//在字典中不能使用容量 只有len,不能使用cap  键值对的个数
	//2。定义方式2
	m2 := make(map [int]string)
	fmt.Println(m2)
	//不能使用cap,但是定义的时候可以指定容量,map可以自动扩容 所以没必要指定容量
	n2 := make(map[int]string ,2)
	fmt.Println(n2)
	//赋值
	n2[1]="张三"
	n2[2]="ls"
	n2[3]="zl"
	n2[4]="zs"
	fmt.Println(n2) //输出的顺序是无序的
	//定义初始化
	m3 := map[int]string{1:"zs",2:"ls",3:"zl"}
	fmt.Println(m3[1])
	//查看字典所在地址
	fmt.Printf("%p\n",m3)
	fmt.Printf("%p\n",m1) //0x0是系统占用 这块空间不允许读写操作
	//4. 使用
	//map中key的类型必须支持 == ！= 一般建议基本数据类型
	//切片 函数 切片的结构类型 不能作为字典的key
	m4 := make(map[string]string)
	//n4:= make(map[string][3]int)//值为数组
	n4:= make(map[string][3]int,1)
	fmt.Println(m4,n4)
	n4["zs"]=[3]int{100,99,98}
	n4["ls"]=[3]int{10,100,90}
	fmt.Println(n4)
	//取值
	for k,v := range n4{
		fmt.Println(k,v)
	}
	fmt.Println(n4)
	fmt.Println(len(n4))
	fmt.Println(n4["ww"]) //找不到值，结果就是string类型的默认值 ""  int的话就是0
	//判断键值对是否存在
	value,ok := n4["zs"]
	fmt.Println(value,ok)  //ok是boolean
	//删除map中的元素，根据键删除
	//如果key不存在，不会报错
	delete(n4,"ls")

	//5，map作为参数传递
	a5 := map[int]string{1:"貂蝉",2:"大乔",3:"小乔"}
	//字典作为参数 也是地址传递 形参和实参指向相同的内存地址 修改形参会影响实参的值
	fmt.Printf("%p\n",a5)
	foo5(a5)
	fmt.Printf("%p\n",a5)
	fmt.Println(a5)
}
func foo5(a5 map[int]string)  {
	a5[4] = "潘金莲"
	fmt.Println(a5)
}