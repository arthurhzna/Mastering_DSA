package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

func (s *Stack) Peek() *Node {
	return s.top
}

func (s *Stack) Push(value string) *Stack {
	newNode := &Node{value: value}
	if s.length == 0 {
		s.top = newNode
		s.bottom = newNode
	} else {
		holdingPointer := s.top
		s.top = newNode
		s.top.next = holdingPointer
	}
	s.length++
	return s
}

func (s *Stack) Pop() *Stack {
	if s.top == nil {
		return nil
	}
	if s.top == s.bottom {
		s.bottom = nil
	}
	s.top = s.top.next
	s.length--
	return s
}

func main() {
	myStack := &Stack{}
	fmt.Println(myStack.Push("google"))
	fmt.Println(myStack.Push("google2"))
	fmt.Println(myStack.Peek())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
}
