package main

import "stack" 
import "fmt"

func main() {
	s := stack.Stack {}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Pop()

	fmt.Println(s)
	
}
