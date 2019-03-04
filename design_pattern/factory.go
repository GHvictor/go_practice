package design_pattern

import (
	"fmt"
)

type Op interface {
	Init()
	Print() string
}
type A struct {}
type B struct {}
type Factory struct {}

func (self A) Init() {
	fmt.Println("init A succ")
}

func (self A) Print() string {
	fmt.Println("init A succ")
	return "A succc"
}

func (self B) Init() {
	fmt.Println("init B succ")
}

func (self B) Print() string {
	return "B succc"
}

func (self Factory) Create(name string) Op {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	switch name {
	case "A":
		return new(A)
	case "B":
		return new(B)
	default:
		panic("name error")

	}
	return nil
}
