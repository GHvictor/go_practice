package leetcode

import "sort"

/*
type ListNode struct {
	Val int
	Next *ListNode
}
*/

func mergeKLists(lists []*ListNode) *ListNode {
	tmp_array := []int{}
	for i := 0; i < len(lists); i++ {
		for lists[i] != nil {
			tmp_array = append(tmp_array, lists[i].Val)
			lists[i] = lists[i].Next
		}
	}
	sort.Ints(tmp_array)

	head := &ListNode{0, nil}
	h := head
	for i := 0; i < len(tmp_array); i++ {
		var new_list *ListNode
		new_list.Val = tmp_array[i]
		new_list.Next = nil
		head.Next = new_list
		head = head.Next
	}
	return h
}
