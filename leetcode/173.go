package leetcode

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	tree []*TreeNode
	num int
}


func Constructor(root *TreeNode) BSTIterator {
	stack := list.New()
	bi := []*TreeNode {}
	cur := root
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			e := stack.Back()
			v := e.Value.(*TreeNode)
			bi = append(bi, v)
			stack.Remove(e)
			cur = v.Right
		}
	}
	return BSTIterator{bi, 0}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.HasNext() {
		ret := this.tree[this.num].Val
		this.num++
		return ret
	}
	return 0
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.tree) <= 0 || this.num >= len(this.tree) {
		return false
	}
	return true
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
