package main

import "fmt"

type Stack struct {
	Elements []string
}

func (s *Stack) Push(item string) {
	s.Elements = append(s.Elements, item)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}

	index := len(s.Elements) - 1
	item := s.Elements[index]
	s.Elements = s.Elements[:index]

	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func main() {
	stack := Stack{}

	stack.Push("Jimmy")
	stack.Push("Tom")
	stack.Push("Alex")

	fmt.Println("Pop item:", stack.Pop())
	fmt.Println("Pop item:", stack.Pop())
	fmt.Println("Pop item:", stack.Pop())

	fmt.Println("Is stack empty?", stack.IsEmpty())
}
