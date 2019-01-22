package leetcode

//tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mergeTree(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	tmp1 := []*TreeNode{t1}
	tmp2 := []*TreeNode{t2}
	for len(tmp1) > 0 && len(tmp2) > 0 {
		tt1 := tmp1[0]
		tt2 := tmp2[0]
		tt1.Val += tt2.Val
		if tt1.Left != nil && tt2.Left != nil {
			tmp1 = append(tmp1, tt1.Left)
			tmp2 = append(tmp2, tt2.Left)
		} else if tt1.Left == nil {
			tt1.Left = tt2.Left
		}
		if tt1.Right != nil && tt2.Right != nil {
			tmp1 = append(tmp1, tt1.Right)
			tmp2 = append(tmp2, tt2.Right)
		} else if t1.Right == nil {
			tt1.Right = tt2.Right
		}
		tmp1 = tmp1[1:]
		tmp2 = tmp2[1:]
	}
	return t1
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil && t2 != nil {
		return t2
	} else if t2 == nil && t1 != nil {
		return t1
	}
	root := new(TreeNode)
	root.Val = t1.Val + t2.Val
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return root
}
