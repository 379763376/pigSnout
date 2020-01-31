package main

import (
	"Bolt_golang_way/uuuuu/tree"
	"fmt"
)

//1.定义别名
type Queue []int

func (q *Queue) Push(v int)  {
	*q = append(*q,v)
}
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

//2.使用组合
/*
元素类型为要扩展的类型指针
如果直接传数据需要把数据拷贝，所以直接传地址

定义自己的结构体方法
对象 != nil
对象.属性 != null
 */
//main 包中的不同的文件的代码不能相互调用，其他包可以。所以其实Link.go没有被一起编译执行
type myTreeNode struct {
	node *tree.Node
}

func main() {
	var boltTree tree.Node
	boltTree = tree.Node{Value:2}
	myTree := myTreeNode{&boltTree}
	fmt.Println(myTree.node.Value)
}
