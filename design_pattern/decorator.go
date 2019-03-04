package design_pattern

import "fmt"

func Double(i int) int {
	return i*2
}

type funObj func(int) int

func LogDecorator(fn funObj) funObj {
	return func(i int) int {
		fmt.Println("start")
		ret := fn(i)
		fmt.Println("finish")
		return ret
	}
}
