package main
<<<<<<< HEAD
=======

>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
import (
	"fmt"
)

type Cat struct {
	Name string
<<<<<<< HEAD
	Age int
=======
	Age  int
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
}

func main() {

	//定义一个存放任意数据类型的管道 3个数据
	//var allChan chan interface{}
	allChan := make(chan interface{}, 3)

<<<<<<< HEAD
	allChan<- 10
	allChan<- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan<- cat
=======
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab

	//我们希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan

	newCat := <-allChan //从管道中取出的Cat是什么?

	//fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	//下面的写法是错误的!编译不通过
	//fmt.Printf("newCat.Name=%v", newCat.Name)
	//使用类型断言
<<<<<<< HEAD
	a := newCat.(Cat) 
	fmt.Printf("newCat.Name=%v", a.Name)
	

}
=======
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)

}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
