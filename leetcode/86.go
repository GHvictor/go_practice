package leetcode

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (self *LinkedList) Append(node *ListNode) {
	current := self.Head
	for {
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	current.Next = node
	self.sizeInc()
}

func (self *LinkedList) sizeInc() {
	self.Head.Val += 1
}

func partition(head *ListNode, x int) *ListNode {
	tmp1 := &ListNode{}
	tmp2 := &ListNode{}
	h_tmp1 := tmp1
	h_tmp2 := tmp2
	for ;head.Next != nil; head = head.Next {
		if head.Val < x {
			tmp1.Next = &ListNode{head.Val, nil}
			tmp1 = tmp1.Next
		} else {
			tmp2.Next = &ListNode{head.Val, nil}
			tmp2 = tmp2.Next
		}
	}
	tmp1.Next = h_tmp2.Next
	return h_tmp1.Next

}

func main() {
	a := "wfewfw"
	fmt.Print(a[1])

}
