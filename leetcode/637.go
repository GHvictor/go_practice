package leetcode

import "container/list"

/**
   3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := list.New()
	queue.PushBack(root)
	ret := []float64{}
	for queue.Len() > 0 {
		cnt := queue.Len()
		sum := 0
		for i := 0; i < cnt; i++ {
			e := queue.Front()
			queue.Remove(e)
			sum += e.Value.(*TreeNode).Val
			if e.Value.(*TreeNode).Left != nil {
				queue.PushBack(e.Value.(*TreeNode).Left)
			}
			if e.Value.(*TreeNode).Right != nil {
				queue.PushBack(e.Value.(*TreeNode).Right)
			}
		}
		ret = append(ret, float64(sum/cnt))
	}
	return ret
}
