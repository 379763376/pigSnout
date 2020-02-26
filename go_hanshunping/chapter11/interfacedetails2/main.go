package main
<<<<<<< HEAD
=======

>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
import (
	"fmt"
)

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

//如果需要实现AInterface,就需要将BInterface CInterface的方法都实现
type Stu struct {
}
<<<<<<< HEAD
=======

>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
func (stu Stu) test01() {

}
func (stu Stu) test02() {
<<<<<<< HEAD
	
}
func (stu Stu) test03() {
	
}

type T  interface{

=======

}
func (stu Stu) test03() {

}

type T interface {
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()

	var t T = stu //ok
	fmt.Println(t)
<<<<<<< HEAD
	var t2 interface{}  = stu
=======
	var t2 interface{} = stu
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t)
<<<<<<< HEAD
}
=======
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
