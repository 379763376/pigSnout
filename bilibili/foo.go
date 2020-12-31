package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	nextNode := addTwoNumbers(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{ Val: sum, Next: nextNode }
	} else {
		tempNode := &ListNode{
			Val: 1,
			Next: nil,
		}
		return &ListNode{
			Val: sum - 10,
			Next: addTwoNumbers(nextNode, tempNode),
		}
	}
}

func main() {
	l1 := new(ListNode)
	l1.Val = 2
	l11 := new(ListNode)
	l11.Val =4
	l111 := new(ListNode)
	l111.Val = 3
	l1.Next = l11
	l11.Next = l111
	l2 := new(ListNode)
	l2.Val =5
	l22 := new(ListNode)
	l22.Val =6
	l222 := new(ListNode)
	l222.Val = 4
	l2.Next = l22
	l22.Next = l222

	fmt.Println(addTwoNumbers(l1,l2).Val)
	fmt.Println(addTwoNumbers(l1,l2).Next.Val)
	fmt.Println(addTwoNumbers(l1,l2).Next.Next.Val)
}