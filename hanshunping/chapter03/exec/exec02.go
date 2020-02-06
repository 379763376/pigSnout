<<<<<<< HEAD
package main
import "fmt"

func main() {

	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int 
	ptr = &num
	*ptr = 10 //这里修改时，会到num的值变化
	fmt.Println("num =" , num)
	var a_b int = 20
	fmt.Println(a_b)

	var int int = 30
	fmt.Println(int)
}
=======
package main

import "fmt"

func main() {

	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int
	ptr = &num
	*ptr = 10 //这里修改时，会到num的值变化
	fmt.Println("num =", num)
	var a_b int = 20
	fmt.Println(a_b)

	var int int = 30
	fmt.Println(int)
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
