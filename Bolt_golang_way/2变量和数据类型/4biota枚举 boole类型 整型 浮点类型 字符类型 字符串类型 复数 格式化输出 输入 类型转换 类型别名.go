/*
内建数据类型

bool string

(u)int (u)int8 （u）int16 (u)int32 (u)int64 uintptr
加u无符号整数 不加u有符号整数
不规定长度的根据系统来32 64
uintptr指针

byte rune
因为char只有一字节，rune是32位 是字符型

float32 float64 complex64 complex128

*/
package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

func main() {
	// 一、常量
	//1.枚举 iota 常量自动生成器，每个一行 自动加1
	//2.iota给常量赋值使用
	const (
		x = iota
		y = iota
		z = iota
	)
	fmt.Println(x, y, z)

	//3.iota遇到const,重置为0
	const d = iota
	fmt.Println(d)
	//4. 可以只写一个iota  在一个括号里
	const (
		a1 = iota
		_
		b1
		c1
	)
	fmt.Println(a1, b1, c1)
	//5. 如果在同一行 ，值是一样的
	const (
		i          = iota             //0
		j1, j2, j3 = iota, iota, iota //1，1，1，
		k          = iota             //2
	)
	const (
		b  = 1 << (10 * iota) //1  iota作为子增值的种子
		kb                    //1024
		mb                    //
		gb
		tb
		pb
	)
	fmt.Println(i)
	fmt.Println(j1, j2, j3)
	fmt.Println(k)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	const aa, bb = 3, 4
	var c2 int
	c2 = int(math.Sqrt(aa*aa + bb*bb))
	print(c2)

	//二、基础数据类型 -- bool 布尔
	//1.bool 布尔类型  占一个字节
	// 声明变量 没有初始化   零值为false
	var a2 bool
	fmt.Println(a2)
	a2 = true
	fmt.Println(a2)
	//2. bool类型不接受其他类型的赋值，不支持强制类型转换
	//a2 = 1
	//3. 自动推到
	var b2 = true
	fmt.Println(b2)
	b3 := false
	fmt.Println(b3)
	b4 := (1 == 2)
	fmt.Println(b4)

	//三、基础数据类型 -- 整型
	var x1 int = 10
	var x2 int = -10
	//1. uint 无符号整型 int 有符号整型  零值为0 占4或8字节（根据操作系统决定）
	var x3 uint = 12
	x4 := 10
	fmt.Println(x1, x2, x3)
	fmt.Printf("%T\n", x4)
	fmt.Println(unsafe.Sizeof(x4)) //查看字节占用大小

	var x5 int32 = 10 //不指定就是int64
	println(x5)
	//四、基础数据类型 -- 浮点
	//单精度浮点型float32 和 双精度浮点型float64 默认为0.0
	//float32 小数点后7位，float64 15位
	var f float64 = 3.14
	//自动推到 创建的变量 默认为float64
	f2 := 3.14
	fmt.Println(f, f2)

	//五、基础数据类型 -- 字符类型
	//1 单引号 字符  双引号 字符串
	var m byte = 'a'
	fmt.Println(m) //输出97，计算机只能识别数字，用asci对应
	//大写字符（A）+32得到小写字符(a)
	//A== 65(asci)  a==97   0==48
	//asci码表0~255
	//uint8  : byte字符类型是uint8的别名    值范围0~255
	//%c 表示占位符号 表示打印输出一个字符
	fmt.Printf("%c", m)    //a
	fmt.Printf("%c", 97)   //a
	fmt.Println("%T\n", m) //得到的是unit8而不是byte
	//\n \t
	fmt.Println('\n') //输出10  \t 输出是9
	// \0的最用是区分字符串和字符（字符串末尾的结束标志） 字符空--asci码0

	//六、基础数据类型 -- 字符串类型
	var str1 string
	str1 = "abc"
	fmt.Println(str1)
	//自动推倒
	str2 := "我爱读书"
	fmt.Printf("%T\n", str2)
	//比较两者
	ch := 'a'
	str3 := "a" //'a''\0'组成的   \0字符串结束标志，打印的时候是打印\0前面的字符
	fmt.Println(ch, str3)
	//len函数 计算字符串中字符的个数
	fmt.Println(len(str1)) //输出3 不包含\0
	fmt.Println(len(str2)) //输出12  go语言一个汉字占3个字符，为了和linux统一处理
	//字符串拼接  +
	str4 := "hello"
	str5 := str4 + str2
	fmt.Println(str5)

	// 字符与字符串区别
	/*
		字符
		1.单引号
		2.字符 一个字符 除了\n \t 注意字符
		字符串
		1.双引号
		2.字符串可以一个或多个字符组成
		3.字符串都是隐藏一个结束符 '\0'
		4.可以操作字符串的每个字符，下标从0开始
	*/
	str6 := "hello go"
	fmt.Println(str6[0], str6[1])         //输出的结果 104 101
	fmt.Printf("%c,$c", str6[0], str6[1]) //输出h,e

	//七、基础数据类型 -- 复数类型
	//complex64 占8字节  两个float32组成
	//complex128 占16字节 两个float64组成
	var v1 complex64
	v1 = 2.1 + 3.14i
	fmt.Println(v1)  //输出 （2.1+3.14i)  加了个括号
	v2 := 2.2 + 3.3i //默认complex128
	fmt.Println(v2)
	//通过内建函数取实部和虚部
	fmt.Println(real(v2), imag(v2))
	v3 := 3 + 4i
	fmt.Println(cmplx.Abs(v3))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) //这种方式的是通用
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         //Exp底数就是1，只用写指数 输出的不是1，而是1.2246467991473532e-16j
	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)    //小数点后三位（0.000+0.000i)
	//八 格式化输出
	/*
		比较重要的几个格式化输出
		%T 值的数据类型
		%c 字符
		%d 整型
		%f 浮点数 默认会有六位小数 不够0补全   %.2f只保留两位小数
		%t true或false输出的布尔值
		%s 字符串
		%v 自动匹配格式输出
	*/
	var f6 = 4.678123789
	var f7 = 4.123
	fmt.Printf("%.2f,%f\n", f6, f7)

	// 九 、输入
	//声明变量 接收用户输入的
	var a int //声明一个变量 在内存中声明一块地址，没有赋值初始化，默认是0
	//格式化输入
	//会阻塞进程 等待用户的输入
	fmt.Printf("请用户输入数据")
	fmt.Scanf("%d", &a) //数据类型 &a是取地址运算符  把用户输入的数据给到a变量
	//输入3.14  打印a  只会输出3

	//更简单的输入方式
	fmt.Scan(&a) //把用户输入的给到变量

	//十、类型转换  (只有强制类型转换）

	//1.三数求和
	m1, m2, m3 := 54, 30, 20
	sum := m1 + m2 + m3
	fmt.Println(sum)
	avg := (sum) / 3
	fmt.Println(avg) //输出34  因为三个原始都是整型 推倒结果也是整型
	//类型转换
	//1.数据类型（变量）int(a)
	//2.数据类型（表达式）
	avg2 := (float64(sum)) / 3
	fmt.Println(avg2)
	//3.布尔不能转其他 其他也不能转bool
	//类型转换是在不同的，但是相互兼容的方式之间转换
	var d32 float32 = 3.1
	var f64 float64 = 3.5
	//在类型转换时 建议从低精度的转成高精度，反则会损失精度  会数据溢出  符号可能会发生改变
	//int8-->int16-->int32-->int64
	//float32-->float64
	//int64 -- float64  建议整型转成浮点型
	num := float64(d32) + f64
	println(num)
	var i64 int64 = 12345
	i8 := int8(i64)
	fmt.Println(i8) //输出57

	//十一、类型别名
	//给int64起一个别名叫bigint
	type bigint int64
	var ab bigint //var ab int64
	fmt.Println(ab)
	//
	type (
		long int64
		char byte
	)
	var bl long = 11
	var cc char = 'a'
	fmt.Println(bl, cc)

}
