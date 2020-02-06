package model
<<<<<<< HEAD
import (
	"fmt"
)
//声明一个Customer结构体，表示一个客户信息 

type Customer struct {
	Id int 
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//使用工厂模式，返回一个Customer的实例 

func NewCustomer(id int, name string, gender string, 
	age int, phone string, email string ) Customer {
	return Customer{
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
=======

import (
	"fmt"
)

//声明一个Customer结构体，表示一个客户信息

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//使用工厂模式，返回一个Customer的实例

func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	}
}

//第二种创建Customer实例方法，不带id
<<<<<<< HEAD
func NewCustomer2(name string, gender string, 
	age int, phone string, email string ) Customer {
	return Customer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
=======
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
	}
}

//返回用户的信息,格式化的字符串
<<<<<<< HEAD
func (this Customer) GetInfo()  string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, 
		this.Name, this.Gender,this.Age, this.Phone, this.Email)
	return info

}
=======
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id,
		this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info

}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
