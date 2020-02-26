<<<<<<< HEAD
package utils
import "fmt"
var Age int
var Name string

//Age 和 Name 全局变量，我们需要在main.go 使用
//但是我们需要初始化Age 和 Name

//init 函数完成初始化工作
func init() {
	fmt.Println("utils 包的  init()...")
	Age = 100
	Name = "tom~"
}

=======
package utils

import "fmt"

var Age int
var Name string

//Age 和 Name 全局变量，我们需要在main.go 使用
//但是我们需要初始化Age 和 Name

//init 函数完成初始化工作
func init() {
	fmt.Println("utils 包的  init()...")
	Age = 100
	Name = "tom~"
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
