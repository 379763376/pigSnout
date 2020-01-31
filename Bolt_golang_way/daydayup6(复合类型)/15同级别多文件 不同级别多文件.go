package main

import (
	"Bolt_golang_way/daydayup6(复合类型)/userInfo"
	"fmt"
)

func main()  {
	//项目中文件管理  要求src（源码） bin（执行文件） pkg(go install 安装的代码包)

	//1. 同级多文件管理  代码在上一级目录src下 main test  同级目录 包名要一致（main）

	//2. 不同级别多文件管理
	// 同级目录 包名要一致
	//包名.函数名  函数名首字母要大写
	fmt.Println()
	userInfo.CreateUser()
}
