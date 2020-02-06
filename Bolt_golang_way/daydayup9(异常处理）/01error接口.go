package main
/*
异常处理;
 */
import (
	"errors"
	"fmt"
)

func test(a, b int) (v int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("不能为0")
	} else {
		v = a / b
	}
	return
}

func main() {

	result,err := test(10, 0)

	if err!= nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}


}
