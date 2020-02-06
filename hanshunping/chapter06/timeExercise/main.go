<<<<<<< HEAD
package main
import (
	"fmt"
	"time"
	"strconv"
)

func test03() {

	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	//在执行test03前，先获取到当前的unix时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03()耗费时间为%v秒\n", end-start)
}
=======
package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {

	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	//在执行test03前，先获取到当前的unix时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03()耗费时间为%v秒\n", end-start)
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
