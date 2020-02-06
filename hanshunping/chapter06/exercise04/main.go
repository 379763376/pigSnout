<<<<<<< HEAD
package main
import (
	"fmt"
)

//请编写一个函数 swap(n1 *int, n2 *int) 可以交换 n1 和 n2的值
func swap(n1 *int, n2 *int) {
	//定义一个临时变量
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	
	a := 10
	b := 20
	swap(&a, &b) //传入的地址
	fmt.Printf("a=%v, b=%v", a, b)
}
=======
package main

import (
	"fmt"
)

//请编写一个函数 swap(n1 *int, n2 *int) 可以交换 n1 和 n2的值
func swap(n1 *int, n2 *int) {
	//定义一个临时变量
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {

	a := 10
	b := 20
	swap(&a, &b) //传入的地址
	fmt.Printf("a=%v, b=%v", a, b)
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
