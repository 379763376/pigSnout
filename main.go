package main

import (
	"fmt"
)

type human interface {
	birthday()
}
type person struct {
	name, id string
}

func (p person) birthday()  {
	fmt.Println("create a human")
}
type stu struct {
	person
	school string
}
type teacher struct {
	person
	tel string
}

func (s stu) learn() string {
	return "good"
}
func (t *teacher) tech() int  {
	return 90
}
func (t *teacher) String() string  {
	return  fmt.Sprintf("xxxx%s",t.tel)
}
func main() {
	var s stu
	var t teacher
	t = teacher{tel: "1111"}
	fmt.Printf("%s",s.learn())
	fmt.Printf("%d",t.tech())
	fmt.Println(&t)
	fmt.Println(t)
	fmt.Println(t.String())
	//fmt.Println(teacher{}.String())

}