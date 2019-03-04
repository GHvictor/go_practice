package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nums1 := []int{}
	nums2 := []int{}

	for ;l1 != nil; l1 = l1.Next {
		nums1 = append(nums1, l1.Val)
	}
	for ;l2 != nil; l2 = l2.Next {
		nums2 = append(nums2, l2.Val)
	}
	count := 0
	head := &ListNode{0, nil}
	for i, j := len(nums1)-1, len(nums2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		a := count
		count = 0
		if i >= 0 && j >= 0 {
			a += nums1[i] + nums2[j]
		} else if i >= 0 && j < 0 {
			a += nums1[i]
		} else if i < 0 && j >= 0 {
			a += nums2[j]
		}

		if a >= 10 {
			count++
			a -= 10
		}

		head.Next = &ListNode{a, head.Next}
	}
	if count == 1 {
		head.Next = &ListNode{count, head.Next}
	}
	return head.Next
}
