package main

import (
	"fmt"
	"github.com/GHvictor/go_practice/design_pattern"
	_ "github.com/GHvictor/go_practice/leetcode"
	"math"
	"reflect"
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

func less (arr []int, i int, j int) bool {
	return arr[i] < arr[j]
}

func swag (arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuiklySort (arr []int) {
	if len(arr) <= 0 {
		return
	}
	k := 0
	flag := true
	for i, j := 0, len(arr) - 1; i < j; {
		if flag {
			if !less(arr, k, j) {
				swag(arr, k, j)
				k = j
				flag = false
			} else {
				j--
			}
		} else {
			if less(arr, k, i) {
				swag(arr, k, i)
				k = i
				flag = true
			} else {
				i++
			}
		}
	}
	QuiklySort(arr[:k])
	QuiklySort(arr[k+1:])
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

func main() {
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
