package main

import (
	"fmt"
	"fun/stack"
)

func main() {
	s := stack.New()

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		fmt.Println("Length should be 2")
	}

	t := s.Pop()
	fmt.Println(t)

	if s.Peek().(int) != 2 {
		fmt.Println("Top of the stack should be 2")
	}
}
