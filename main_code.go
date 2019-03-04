package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"github.com/GHvictor/go_practice/design_pattern"
	_ "github.com/GHvictor/go_practice/leetcode"
	"math"
	"reflect"
	"strings"
)

type Phone interface {
	call()
}

type AndroidPhone struct {
}

func (self AndroidPhone) call() {
	fmt.Println("I am Android, I can call you!")
}

type IPhone struct {
}

func (self IPhone) call()  {
	fmt.Println("I am iPhone, I can call you!")
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (self Rectangle) area() float64 {
	return self.width * self.height
}

func (self Circle) area() float64 {
	return self.radius * self.radius * math.Pi
}

func less(i int, j int) bool {
	return i < j
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSort(arr []int) {
	if len(arr) < 1 {
		return
	}
	mid := arr[0]
	i := 0
	head, tail := 0, len(arr)-1
	flag := true
	for head < tail {
		if flag {
			if less(arr[tail], mid) {
				swap(arr, i, tail)
				i = tail
				flag = false
			} else {
				tail--
			}
		} else {
			if less(mid, arr[head]) {
				swap(arr, i, head)
				i = head
				flag = true
			} else {
				head++
			}

		}
	}
	QuickSort(arr[:i])
	QuickSort(arr[i+1:])

}

type sn1 struct {
	name string
	age int
}

type sn2 struct {
	name string
	age int
}
type User struct {
}
type MyUser1 User
type MyUser2 = User
func (i MyUser1) m1(){
	fmt.Println("MyUser1.m1")
}
func (i MyUser1) m2()  {
	fmt.Println("MyUser")

}
func (i User) m2(){
	fmt.Println("User.m2")
}

func test_code() {
	a := list.New()
	sn := &sn2{"z", 18}
	a.PushBack(sn)
	q := a.Back()
	q.Prev()
}

type IntHeap []int  // 定义一个类型

func (h IntHeap) Len() int { return len(h) }  // 绑定len方法,返回长度
func (h IntHeap) Less(i, j int) bool {  // 绑定less方法
	return h[i] > h[j]  // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntHeap) Swap(i, j int) {  // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {  // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {  // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func main() {
	var uaaa, ubbb, uccc uint64
	uaaa = 4
	ubbb = 8
	uccc = uaaa ^ ubbb
	fmt.Println(uccc)
	count := 0
	fmt.Println(uccc&uint64(1))
	for i := 0; i < 64; i++ {
		if uccc&uint64(1)  == uint64(1) {
			count++
		}
		uccc>>=1
	}
	fmt.Println(count)

	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}  // 创建slice
	heap.Init(h)  // 初始化heap
	fmt.Println(*h)
	qh := *h
	fmt.Println(qh[0])
	fmt.Println(qh[len(qh)-1])
	fmt.Println(heap.Pop(h))  // 调用pop
	/*
	heap.Push(h, 6)  // 调用push
	fmt.Println(*h)
	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}*/

	qc := "ccccc,dcoewj"
	dc := strings.Split(qc, ",")
	qqq := list.New()
	qqq.PushBack("fejwioewf")
	m := qqq.Back()
	if ii, ok := m.Value.(string); ok {
		fmt.Println(ii)
	} else {
		fmt.Println("jfoweoewfojoweoewfoew")
	}
	fmt.Println(reflect.TypeOf(qc))
	for _, v := range dc {
		fmt.Println(v)
	}
	instance := design_pattern.GetInstance()
	instance.Init()

	f := design_pattern.Factory{}
	a := f.Create("A")
	a.Init()
	fmt.Println(a.Print())
	b := f.Create("B")
	b.Init()
	fmt.Println(b.Print())
	c := f.Create("C")
	aValue := reflect.ValueOf(a)
	fmt.Println(aValue.Type())
	bValue := reflect.ValueOf(b)
	fmt.Println(bValue.Type())
	fmt.Println(c)
	ld := design_pattern.LogDecorator(design_pattern.Double)
	fmt.Println(ld(44))

}
