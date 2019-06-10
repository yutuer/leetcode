package main

import "fmt"
import "github.com/yutuer/mydatastruct/stack"

func main() {
	fmt.Println("test")

	myStack := stack.NewStack()
	myStack.Push(1)

	i, _ := myStack.Pop()

	fmt.Println(i)
}
