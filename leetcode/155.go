package leetcode

import (
	"container/list"
	"math"
)

type MinStack struct {
	min int
	stack *list.List
}


/** initialize your data structure here. */
func Constructorl() MinStack {
	return MinStack{
		nil,
		list.New(),
	}
}


func (this *MinStack) Push(x int)  {
	if this.min > x {
		this.min = x
	}
	this.stack.PushBack(x)
}


func (this *MinStack) Pop()  {
	if this.stack.Len() > 0 {
		this.stack.Remove(this.stack.Back())
	}
	min := math.MaxInt64
	for i := this.stack.Front(); i != nil; i = i.Next() {
		if min > i.Value.(int) {
			min = i.Value.(int)
		}
	}
	this.min = min
}


func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}


func (this *MinStack) GetMin() int {
	return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
