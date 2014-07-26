go-stack
========
go-stack implements a singly linked list stack queue

Installation
============
To get the package, type:
`go get github.com/paulspychala/go-stack

Example
=======
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

Contributors
============
* [Pawel Spychala](https://github.com/paulspychala)
