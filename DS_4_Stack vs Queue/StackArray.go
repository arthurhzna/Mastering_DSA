package main

import "fmt"

type Stack struct {
	array []string
}

func (s *Stack) Peek() string {
	if len(s.array) == 0 {
		return ""
	}
	return s.array[len(s.array)-1]
}

func (s *Stack) Push(value string) *Stack {
	s.array = append(s.array, value)
	return s
}

func (s *Stack) Pop() *Stack {
	if len(s.array) > 0 {
		s.array = s.array[:len(s.array)-1]
	}
	return s
}

func main() {
	myStack := &Stack{}
	myStack.Peek()
	myStack.Push("google")
	myStack.Push("ztm")
	myStack.Push("discord")
	fmt.Println(myStack.Peek())
	myStack.Pop()
	myStack.Pop()
	myStack.Pop()
}
