package main

import "fmt"

/*
结构体
定义 在函数外部
	type 结构体名 struct{
		结构体成员列表
		成员名 数据类型
		name string
		age int
	}
结构体名可见性   结构名和函数体在其他文件也可以被导入 就要求名字首字母大写
 */
//结构体名 大小写
type student struct {
	id int
	name string
	sex string
	age int
	addr string
}
type Student struct {
	id int
	name string
	score int
}
func main() {
	//使用结构体
	//顺序初始化 每个成员必须初始化
	var s1 = student{1,"黎明", "男",23,"社区"}
	fmt.Println(s1)

	//自动推导类型 和 指定成员赋值
	s2 := student{name:"小娟",age:18,sex:"女"}
	fmt.Println(s2)

	//定义结构体变量  复合类型
	//var 变量名 结构体名
	var s3 student
	//赋值
	//为结构体成员赋值  ，包名.函数  结构体.成员  对象.方法
	s3.name="小花"
	fmt.Println(s3)
	fmt.Printf("%p\n",&s3)  //结构体查看内存地址需要有&修饰符
	fmt.Printf("%p\n",&s3.id)  //结构体名指向第一个成员的地址 结果体在内存中占用的大小为成员占用的总和
	fmt.Printf("%p\n",&s3.name)
	//结构体赋值和比较
	a3 := student{1,"小张", "男",23,"社区"}
	b3 := student{2,"小王", "男",23,"社区2"}
	c3 := student{id:3,name:"小李",sex:"女"}
	fmt.Println(a3 == b3)  //结构体的比较
	b3 = a3  //同类型的结构体 可以相互赋值
	c3 =a3
	fmt.Println(a3 == b3)
	fmt.Println(c3 == b3)

	//结构体数组 结构体切片
	var arr [3]Student = [3]Student{
		Student{id:1,name:"zs",score:90},
		Student{id:2,name:"ls",score:80},
		Student{id:3,name:"zl",score:89}}
	fmt.Println(arr)
	fmt.Println(arr[0])
	//for循环打印结构体
	//修改结构体中的信息
	arr[0].score = 92
	fmt.Println(arr)
	//结构体数组传递也是值传递

	
}

func Sort(ars[3]Student){
	for i :=0 ;i<len(ars)-1;i++{//比较了多少次
		for j:=0;j<len(ars)-i-1;j++{ //每趟对比的次数
			if ars[j].score>ars[j+1].score{
				ars[j],ars[j+1] = ars[j+1],ars[j] //大的往后移
			}
		}
	}
}