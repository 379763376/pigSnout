<<<<<<< HEAD
package main
import (
	"fmt"
	"unsafe"
)

//演示golang中bool类型使用
func main() {
	var b = false
	fmt.Println("b=", b)
	//注意事项
	//1. bool类型占用存储空间是1个字节
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b) )
	//2. bool类型只能取true或者false
	
}
=======
package main

import (
	"fmt"
	"unsafe"
)

//演示golang中bool类型使用
func main() {
	var b = false
	fmt.Println("b=", b)
	//注意事项
	//1. bool类型占用存储空间是1个字节
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
	//2. bool类型只能取true或者false

}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
