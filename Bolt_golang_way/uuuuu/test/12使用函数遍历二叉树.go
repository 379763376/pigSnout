package main

//递归遍历的时候不一定只是打印 需要做很多事情 之前的print有些简陋
type MyTree struct {
	Value       int
	Left, Right *MyTree
}
func (node *MyTree) MyTraverse() {
	if node == nil {
		return
	}
	node.Left.MyTraverse ()
	node.MyPrint()
	node.Right.MyTraverse()
}
func (node MyTree) MyPrint() {
	node.MyTraverseFunc(func(n *MyTree) {
		n.MyPrint()
	})
}
func (node *MyTree) MyTraverseFunc(f func(*MyTree)) {
	if node == nil {
		return
	}
	node.Left.MyTraverseFunc (f)
	f(node)
	node.Right.MyTraverseFunc(f)
}