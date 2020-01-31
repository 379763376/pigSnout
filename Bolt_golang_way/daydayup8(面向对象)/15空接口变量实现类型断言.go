package main

import (
	"fmt"
)

func main1501()  {
	var i interface{}
	i = 10.34
	//value, ok := element.(T)
	//值，值的判断 := 接口变量.（数据类型）

	value,ok := i.(int)
	if ok {
		fmt.Println("整型数据",value)
	}else {
		fmt.Println("不是整型")
	}

}
func demo15()  {
	fmt.Println("hello")
}

//类型断言
func main1502()  {
	var i []interface{}
	i = append(i, 10,3.14,"aaa",demo15)
	//fmt.Println(i)
	for _,v := range i{
		//fmt.Println(v)
		if data,ok := v.(int);ok {
			fmt.Println("整型数据",data)
		} else if data,ok:=v.(float64);ok {
			fmt.Println("浮点型数据",data)
		}else if data,ok:=v.(string);ok {
			fmt.Println("字符串",data)
		}else if data,ok:=v.(func());ok {
			//函数调用
			data()
		}
	}
}
//switch测试
func main()  {
	var i []interface{}
	i = append(i, 10,3.14,"aaa",demo15)
	for _,v:=range i{
		//value, ok := element.(T)
		switch value:= v.(type) {
		case int:
			fmt.Println("整型数据",value)
		case float64:
			fmt.Println("浮点型数据",value)
		case string:
			fmt.Println("字符串",value)
		case func():
			value()
		}
	}

}