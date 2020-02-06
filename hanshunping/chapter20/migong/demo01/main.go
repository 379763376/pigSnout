<<<<<<< HEAD
package main
import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n-- // 死龟
		test(n)
	} else {
		fmt.Println("n=", n)
	}
	
}

func main() {

	n := 4
	test(n)
}
=======
package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n-- // 死龟
		test(n)
	} else {
		fmt.Println("n=", n)
	}

}

func main() {

	n := 4
	test(n)
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
