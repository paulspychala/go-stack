// Copyright 2014 Pawel Spychala

package main

import (
	"fmt"
	"github.com/paulspychala/go-stack"
)

func main() {
	stack := stack.New()
	stack.Push(3)
	stack.Push(5)
	stack.Push(8)

	for e := stack.Pop(); e != nil; e = stack.Pop() {
		fmt.Println(e)
	}
}
