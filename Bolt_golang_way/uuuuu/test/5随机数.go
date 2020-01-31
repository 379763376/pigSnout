package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机数生成 种子+算法 ==> 生成的数就是固定的
	rand.Seed(2)
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(33)) //0~32
	fmt.Println(rand.Intn(33))
	rand.Seed(time.Now().UnixNano())//利用事件生成随机数种子 纳秒

}
