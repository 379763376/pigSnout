package main

import "fmt"

func main()  {
	//数组 同一类型数据的集合
	//1. 定义   var 数组名 [长度] 类型
 	var a[10] int
 	//数组长度
 	fmt.Println(len(a))
	//长度必须制定 固定个一个整数
	//初始化，默认值是0
	//给数组赋值 按照下标赋值 0~9
	a[0] = 1
	a[1] = 2
	fmt.Println(a)
	fmt.Println(a[9])
	// fmt.Println(a[10])
	//使用for循环赋值
	//使用for循环取出每个元素
	//for _,v := range a 打印输出数组元素
	//定义float64 默认0
	//var b[10]float64
	//定义string  默认为空
	//var c[10]string
	//var d[10]bool //默认false
	//2. 初始化
	var e [5]int = [5]int{1,2,3,4,5}
	fmt.Println(e)
	// 自动推导
	f := [5]int{1,2,3,4,5}
	fmt.Println(f)
	// 部分初始化  没有初始化的赋值为0
	g:= [5]int{1,2,3}
	fmt.Println(g)
	// 指定某个元素初始化
	h:= [5]int{2:10,4:20}
	fmt.Println(h)
	// ...通过初始化确认数组长度
	i:=[...]int{1,2,3}
	fmt.Println(len(i))
	//3. 常见问题
	//数组长度 必须是一个常量 不能是变量
	//下标  不能导致数组下标越界
	//i[-1] go中-1不表示最后一个元素
	//如果两个数组元素个数相同 类型相同 是可以赋值的
	j:=i
	fmt.Println(j)
	// 查看数组类型
	fmt.Printf("%T\n",j)
	//数组名表示整个数组  数组名对应的地址就是数组第一个元素的地址
	fmt.Printf("%p\n",&j)
	fmt.Printf("%p\n",&j[0])
	fmt.Printf("%p\n",&j[1])
	fmt.Printf("%p\n",&j[2])

	//4. 实战 数组中取出最大 最小 总和 平均数
	var k [6]int = [6]int{1,2,3,4,5,6}  //{-1，-2，-3，-4，-5，-6}
	var max int = k[0]
	var min int = k[0]
	var sum int
	var avg float64
	for i:=0;i<len(k) ;i++  {
		if k[i]>max{
			max = k[i]
		}
		if k[i]<min{
			min = k[i]
		}
		sum += k[i]
	}
	avg = float64(sum)/float64(len(k))
	fmt.Println(max,min,sum,avg)//最后输出最小值为0  所以需要先给变量赋值为元素的起始值

	//5。 数组的逆置
	var arr[10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	//for 表达式1；表达式2；表达式3
	//for 返回值 ：= range 集合{}
	// for 条件{}
	i1 := 0
	j1:=len(arr)-1
	for i1<j1{
		//if i1>=j1{
		//	break
		//}
		arr[i1],arr[j1]=arr[j1],arr[i1]
		i1++
		j1--
	}
	fmt.Println(arr)

	//6. 数组冒泡排序 从小到大
	var ars[10]int = [10]int{3,1,2,7,5,6,4,10,9,8}
	//{1,2,3,5,6,4,7，9，8，10} 没执行一趟确定一个数据的位置  执行次数9次 i=len(ars)-1次
	//10个元素走9趟就能确定了
	//i表示总共需要执行多少周 i= len(ars)-1
	//j 交换了多少次  第一趟9次 第二趟8次  。。。 第九趟1次

	//外层控制行
	for i :=0 ;i<len(ars)-1;i++{//比较了多少次
		//内层控制列
		for j:=0;j<len(ars)-i-1;j++{ //每趟对比的次数
			//满足条件进行交换 大于升序 小于降序
			if ars[j]>ars[j+1]{
				ars[j],ars[j+1] = ars[j+1],ars[j] //大的往后移
			}
		}
	}
	fmt.Println(ars)
}
