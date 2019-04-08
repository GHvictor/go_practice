package leetcode

/**
Input: 1->2
Output: false
Input: 1->2->2->1
Output: true
 */

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	reverse(slow)
	for head != nil && slow != nil {
		if head.Val == slow.Val {
			head = head.Next
			slow = slow.Next
		} else {
			return false
		}
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newln := head.Next
	newh := reverse(newln)
	newln.Next = head
	head.Next = nil
	return newh
}
