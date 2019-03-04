package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
[1,2,2,3,4,4,3]
    1
   / \
  2   2
 / \ / \
3  4 4  3
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isCheck(root.Left, root.Right)
}

func isCheck(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left != nil && right == nil) || (left == nil && right != nil) || left.Val != right.Val {
		return false
	}
	return isCheck(left.Left, right.Right) && isCheck(left.Right, right.Left)
}
