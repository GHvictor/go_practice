package design

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton{
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (self *singleton) Init() {
	fmt.Println(&self, "init succ")
}
