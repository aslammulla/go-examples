package main

import (
	"fmt"
	"leetcode-go/Aslam/Go-Projects/go-examples/generics/stack"
)

func main() {
	// Stack of integers
	intStack := stack.Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intVal, ok := intStack.Pop()
	fmt.Println("Popped:", intVal, ok)

	top, ok := intStack.Peek()
	fmt.Println("Peek:", top, ok)               // Output: 1, true
	fmt.Println("Size:", intStack.Size())       // Output: 1
	fmt.Println("IsEmpty:", intStack.IsEmpty()) // Output: false

	// Clear the integer stack
	intStack.Clear()
	fmt.Println("After Clear - Size:", intStack.Size())       // Output: 0
	fmt.Println("After Clear - IsEmpty:", intStack.IsEmpty()) // Output: true

	// Stack of strings
	strStack := stack.Stack[string]{}
	strStack.Push("hello")
	strStack.Push("world")

	strVal, ok := strStack.Pop()
	fmt.Println("Popped:", strVal, ok)

	topStr, ok := strStack.Peek()
	fmt.Println("Peek:", topStr, ok)            // Output: hello, true
	fmt.Println("Size:", strStack.Size())       // Output: 1
	fmt.Println("IsEmpty:", strStack.IsEmpty()) // Output: false

	// Clear the string stack
	strStack.Clear()
	fmt.Println("After Clear - Size:", strStack.Size())       // Output: 0
	fmt.Println("After Clear - IsEmpty:", strStack.IsEmpty()) // Output: true
}
