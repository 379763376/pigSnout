package main

import (
	"fmt"
)
//1.结构体
//不论地址还是结构体本身，一律使用“.访问成员
type TreeNode struct {
	value       int
	left, right *TreeNode
}

func main() {
	var root TreeNode
	root = TreeNode{value: 3}
	root.left=&TreeNode{}
	root.right=&TreeNode{5,nil,nil}
	root.right.left=new(TreeNode)

	fmt.Println(root.value)
	fmt.Println(root.left.left) //返回nil
	fmt.Println(root.left.value) //返回zero value
	fmt.Println(root.right.value)
	fmt.Println(root.right.left)
	fmt.Println(root.right.right)


	//2.struct提供了构造方法，不再需要构造函数，可以加工厂函数
	//工厂函数返回指针类型 内部是构造方法
	//方法返回的是一个局部变量给主方法（go语言特有 编译器根据程序是否返回局部变量地址决定在堆栈创建 ）
	var tree BoltTree
	tree = BoltTree{Value:5}
	tree.Left = CreateNode(2) //工厂函数生成节点
	tree.Right = CreateNode(3)
	fmt.Println(tree.Left)
	tree.Left.Left = CreateNode(3)
	fmt.Println(tree.Left.Left)
	tree.Left.Left.SetValue(4) //指针接收者结构体方法  修改结点值
	fmt.Println(tree.Left.Left)
	tree.Left.Left.Print()  //值接收者结构体方法 调用结构体方法

	//6.遍历树
	//中序遍历
	fmt.Println("中序遍历")
	tree.Traverse()
}
type BoltTree struct {
	Value       int
	Left, Right *BoltTree
}
//3.工厂函数
func CreateNode(value int) *BoltTree {
	return &BoltTree{Value: value}
}
//4.结构体方法 指针接收者
func (node *BoltTree) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}
//5.值接收者
func (node BoltTree) Print() {
	fmt.Print(node.Value, " ")
}
//使用递归的思想
func (node *BoltTree) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse ()
	node.Print()
	node.Right.Traverse()
}