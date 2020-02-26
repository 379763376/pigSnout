<<<<<<< HEAD
package main
import (
	"fmt" 
)
func main() {

	//有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
	var a int = 10
	var b int = 20

	a = a + b //
	b = a - b // b = a + b - b ==> b = a
	a = a - b // a = a + b - a ==> a = b

	fmt.Printf("a=%v b=%v", a, b)
}
=======
package main

import (
	"fmt"
)

func main() {

	//有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
	var a int = 10
	var b int = 20

	a = a + b //
	b = a - b // b = a + b - b ==> b = a
	a = a - b // a = a + b - a ==> a = b

	fmt.Printf("a=%v b=%v", a, b)
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
