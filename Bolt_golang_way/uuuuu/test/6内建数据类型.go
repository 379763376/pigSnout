package main

import (
	"fmt"
	"sort"
)

func main() {
	//1.数组
	arr := [...]int{1, 2, 3}
	fmt.Println(arr)
	/*
	   冒泡排序
	   总共执行len(arr)-1趟
	   每趟比较的个数len(arr)-第几趟
	*/
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if(arr[j]>arr[j+1]){
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	//2.slice
	/*
	slice是arr的视图，对slice的修改会改变arr
	ptr指向slice的开头元素在array的地址
	len(slice）是视图的长度
	cap(slice) 是ptr开始到array的长度
	当slice是基于array reslice生成时，且len(slice)<cap(slice) 可以把数据通过reslice视图的方式取出，
	但是不能用下标取出大于len(slice)-1的数据

	cap 每次放不下了长度就会在原来的基础上扩大二倍
	当容量增加到1024以后  扩容上次cap的1/4

	make 创建slice map chan  分配内存 初始化对象 返回对象的引用，区别与new返回对象的地址
	 */
	arr2 :=[...]int{0,1,2,3,4,5,6,7}
	//arr3 := arr2[:]
	//fmt.Println(arr3)
	s2:=arr2[5:6]
	s3 := s2[0:2]
	fmt.Println(s3)
	s4 := append(s3, 11)
	fmt.Println("arr2:",arr2)
	fmt.Println(s4)
	s5:=append(s4,10)
	fmt.Println("arr2:",arr2)
	fmt.Println("s5:",s5)
	//拷贝 定义s6不能时var ，需要返回s6的引用
	s6 := make([]int,len(s5))
	copy(s6,s5)
	//删除  通过reslice
	//排序
	sort.Ints(s6)
	fmt.Println(s6)
	//3.map
	/*
	map的key使用hashmap
	key必须可以比较相等（可以hash）,所以key不能是slice map function,如果struct不含上述类型 也可以作为key
	 */
	//遍历
	mmp := make(map [string]string)
	for k,v := range mmp{
		fmt.Println(k,v)
	}
	v :=mmp["k"] //不存在会打印zero value
	fmt.Println(v)
	if v,ok := mmp["k"];ok{
		fmt.Println(v)
	}

}
