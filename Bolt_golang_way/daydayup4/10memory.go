package main

func main() {
	// 内存   ---   堆栈
	//从主函数运行
	//加载函数信息和函数变量名到栈区内存
	//栈区内存的存储模型
	//内存相对一个可执行程序来说  可以分为 ：代码区 数据区 堆区 栈区
	//栈区： 用来存储程序执行过程中函数内部定义的信息 和 局部变量
	//形式参数  告知调用函数时的格式
	//实参 在调用函数的时候将实参传递给形参
	//函数执行调用完毕就会从内存中释放  栈区存储原理：先进后出，函数先销毁是后加载进来的


	//2. 内存 函数返回值
	//函数调用执行完后会把结果赋值到一个新的变量 ，函数自身销毁了

	//3. 内存 函数的嵌套调用

}
func add(s1,s2 int)  {

}