<<<<<<< HEAD
package main
import (
	"fmt"
)

func test(slice []int) {
	slice[0] = 100  //这里修改slice[0],会改变实参
	}
	
func main() {

	var slice = []int {1, 2, 3, 4}
	fmt.Println("slice=", slice) // [1,2,3,4]
	test(slice)
	fmt.Println("slice=", slice) // [100, 2, 3, 4]


}
=======
package main

import (
	"fmt"
)

func test(slice []int) {
	slice[0] = 100 //这里修改slice[0],会改变实参
}

func main() {

	var slice = []int{1, 2, 3, 4}
	fmt.Println("slice=", slice) // [1,2,3,4]
	test(slice)
	fmt.Println("slice=", slice) // [100, 2, 3, 4]

}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
