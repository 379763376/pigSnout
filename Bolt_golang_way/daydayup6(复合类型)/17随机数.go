package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1.导入头文件 math/rand
	//2.随机数种子
	//3.创建随机数

	//创建随机数种子
	rand.Seed(1)
	fmt.Println(rand.Int())
	fmt.Println(rand.Int()) //使用int生成较大的随机数，
	fmt.Println(rand.Int())
	// 再次运行发现没有变化==》随机数是根据随机数种子计算出来的一个数（计算方法一定 随机数就不会变化）

	//利用事件生成随机数种子
	rand.Seed(time.Now().UnixNano()) //纳秒
	fmt.Println(rand.Int())

	//生成一百以内的数字
	fmt.Println(rand.Intn(100))  //取模100  得到0~99

	//模拟双色球
	// 红球1-33 选6
	// 蓝球1-16 选1
	// 蓝球红球数字可重复
	rand.Seed(time.Now().UnixNano())
	var red [6] int
	for i:=0;i<len(red) ;i++  {
		v := rand.Intn(33)+1  //因为值是0-32 ,所以加1
		for j:=0;j<i ;j++  { //把这个新得到的v和前面的数字比较，查看是否有重复值
			if red[j] == v {
				v =rand.Intn(33)+1
				j -- //把j回退到当前位置 重新比较
			}
		}
		red[i] = v
	}
	fmt.Println("红球",red,"蓝球",rand.Intn(16)+1)

}