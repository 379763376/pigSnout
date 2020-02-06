package main

import "fmt"

/*
1.方法一：类似于排序的比较 两层for嵌套
2.方法二:
x x a b c d x x a b d e
把字符串转换成切片range循环
先以第一个字符x作为start，x之前没出现过，maxLength变为1,把x的位置放入lastOccurred
第二个x x之前没出现过 lastIndex == start
循环到d maxLength变为5，lastOccurred中有x a b c d和其对应的位置
当再到x lastOccurred中有x 把start位置置为当前的x,maxLength不变，lastOccurred更新为当前的位置
直到最后一个字符 得到maxLength

 */
func maxNorepeatSubstring(s string) (maxStart,maxLength int) {
	lastOccurred := make(map[rune]int)
	start := 0
	for i,ch:=range[]rune(s){
		//0x start指向第0个元素
		//1x start指向第1个元素
		//2a 3b 4c 5d
		//6x start指向上一个x后面的元素 第2个元素
		//7x start指向上一个x后面的元素 第6个元素
		//8a  出现过但是比start小，所以没用 所以要加限制条件v > start
		//9b 10d 11e
		if v,ok:=lastOccurred[ch];ok && v >= start {//在字典中出现过
			start = v+1
		}
		if i-start+1>maxLength{
			maxLength = i-start+1
			maxStart = start
		}
		//放入字典
		lastOccurred[ch] = i
	}
	return maxStart,maxLength

}
func main() {
	fmt.Println(
		maxNorepeatSubstring("abcabcbb"))
	fmt.Println(
		maxNorepeatSubstring("bbbbb"))
	fmt.Println(
		maxNorepeatSubstring("pwwkew"))
	fmt.Println(
		maxNorepeatSubstring(""))
	fmt.Println(
		maxNorepeatSubstring("b"))
	fmt.Println(
		maxNorepeatSubstring("abcdef"))
	fmt.Println(
		maxNorepeatSubstring("这里是慕课网"))
	fmt.Println(
		maxNorepeatSubstring("一二三二一"))
	fmt.Println(
		maxNorepeatSubstring(
			"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
