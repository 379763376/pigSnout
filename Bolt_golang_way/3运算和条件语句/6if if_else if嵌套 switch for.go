package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//一 If选择结构
	//1.
	a1 := 700
	var b1 int
	//fmt.Printf("请用户输入分数")
	//fmt.Scanf("%d", &b1)
	//if b1 >= a1 {
	//	fmt.Println("恭喜你1")
	//}
	if b1 >= a1 {
		fmt.Println("恭喜你2")
	}
	if c1 := 710; c1 > a1 { //支持一个初始化语句，初始化语句和判断条件用 ； 分割
		fmt.Println("恭喜你3")
	}
	// 2 if else
	if c2 := 650; c2 > a1 { //支持一个初始化语句，初始化语句和判断条件用 ； 分割
		fmt.Println("恭喜你4")
	} else { //没有条件
		fmt.Println("需要再努力")
	}
	// 3 if elseif else
	a2 := 650
	if c2 := 650; c2 >= a1 { //支持一个初始化语句，初始化语句和判断条件用 ； 分割
		fmt.Println("恭喜你一本")
	} else if c2 >= a2 {
		fmt.Println("恭喜你二本")
	} else {
		fmt.Println("需要再努力")
	}
	/* 4 if和if_else区别
	if 都会去执行
	if_else效率高,只要一个匹配到之后的不再执行
	*/
	/* 5 if嵌套
	if condition1 {
		if condition2 {
			//todo
		}
	}
	*/
	//6，比较的时候不能  x<y<z    z<y结果未boole <z的时候就会报错
	const filename = "abc.txt"
	if contents,err := ioutil.ReadFile(filename);err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}

	//二 switch
	/*  1.支持多个条件的匹配
	a.不同的case之间不需要使用break
	b. fallthrough  强制往下执行一个case
	switch 变量（或表达式） {
		case val1:
			...
		case val2:
			...
		default:
			...
	}
	*/
	score := 90
	switch score {
	case 90:
		fmt.Println('C')
	case 80:
		fmt.Println("B")
	case 10, 20, 30, 40:
		fmt.Println("D")
	default:
		fmt.Println("A")
	}
	//2
	switch s1:=90;s1 {  //初始化语句   ;    条件
	case 90:
		fmt.Println('C')
	case 80:
		fmt.Println("B")
	case 10, 20, 30, 40:
		fmt.Println("D")
	default:
		fmt.Println("A")
	}
	//3
	var s2 int = 90
	switch  { //没有条件
	case s2>90://写判断语句
		fmt.Println('A')
	}
	//4
	switch s3:=90; { //只有初始化没有条件
	case s3>90:
		fmt.Println('A')
	}
	/*if和switch的选择：
		if适合区间判断，嵌套使用
		switch执行效率高，将多个满足条件的值放一起，适合做固定值的判断

		if相对switch执行效率低，分支少用if
		switch不建议使用嵌套
	 */

	/* 循环
		for
	 */
	//1.语法
	//for 初始化条件;判断条件 ;条件变化  {
	//
	//}
	//计算1+2+3+。。。+10的和
	sum := 0
	for i:=1;i<=10;i++{
		sum+=i
	}
	fmt.Println(sum)
	str := "abc"
	for i:=1 ; i<len(str);i++{
		fmt.Println(str[i])
	}
	//迭代打印每个元素 默认返回两个值 第一个元素位置 一个元素本身
	for i,data :=range str  {
		fmt.Println(i,data)
	}
	for i,_ := range str{
		fmt.Println(i)
	}
	for _,data := range str{
		fmt.Println(data)
	}
	for i := range str{
		fmt.Println(i)
	}
	//for 嵌套
	count := 0
	for i :=0;i<5 ;i++  {
		for j :=0;j<5 ;j++  {
			fmt.Println(i,j)
			count +=1
		}
	}
	fmt.Println(count)

	//百钱百鸡
	// 5cock 3hen 1/3chicken
	//方法一：
	count1 := 0
	for cock:=0;cock<=20;cock++{
		for hen:=0;hen<=33 ;hen++  {
			for chicken:=0;chicken<=100 ;chicken+=3  {
				count1+=1
				if 5*cock+3*hen+chicken/3 == 100 &&cock+hen+chicken==100 {
					fmt.Println(cock,hen,chicken)
				}
			}
		}
	}
	//方法二：
	count2 := 0
	for cock:=0;cock<=20;cock++{
		for hen:=0;hen<=33 ;hen++  {
			count2 +=1
			chicken := 100-cock-hen
			if chicken % 3 ==0 && 5*cock+3*hen+chicken/3 == 100{
				fmt.Println(cock,hen,chicken)
			}
		}
	}
	fmt.Println(count1,count2)

}
//for
// 整数转二进制
// 省略初始条件
//省略所有的死循环

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//func forever() {
//	for {
//		fmt.Println("abc")
//	}
//}

func main01() {
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)

	fmt.Println("abc.txt contents:")
	printFile("basic/branch/abc.txt")

	fmt.Println("printing a string:")
	s := `abc"d"
	kkkk
	123

	p`
	printFileContents(strings.NewReader(s))

	// Uncomment to see it runs forever
	// forever()
}