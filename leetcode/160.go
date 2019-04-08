package leetcode


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
Output: Reference of the node with value = 8
Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5].
There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointA, pointB := headA, headB
	if headA == nil || headB == nil {
		return nil
	}
	markA, markB := 0, 0
	for pointA != pointB {
		if pointA == nil {
			if markA >= 1 {
				return nil
			}
			pointA = headB
			markA++
		} else if pointB == nil {
			if markB >= 1 {
				return nil
			}
			pointB = headA
			markB++
		} else {
			pointA = pointA.Next
			pointB = pointB.Next
		}
	}
	return pointA
}