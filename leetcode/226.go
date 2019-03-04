package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func changeTree(left *TreeNode, right *TreeNode) {
	if left == nil && right == nil {
		return
	} else if left != nil && right != nil {
		left.Val, right.Val = right.Val, left.Val
		changeTree(left.Left, right.Right)
		changeTree(left.Right, right.Left)
	} else if left == nil && right != nil {
		left = right
		right = nil
		changeTree(left.Left, left.Right)
	} else if left != nil && right == nil {
		right = left
		fmt.Println("hahaha")
		left = nil
		changeTree(right.Left, right.Right)
	}

}
