package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkedList(value int) *LinkedList {
	node := &Node{Value: value}
	return &LinkedList{Head: node, Tail: node, Length: 1}
}

func (l *LinkedList) Append(value int) *LinkedList {
	newNode := &Node{Value: value}
	if l.Length == 0 {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Length++
	return l
}

func (l *LinkedList) Prepend(value int) *LinkedList {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
	if l.Length == 0 {
		l.Tail = newNode
	}
	l.Length++
	return l
}

func (l *LinkedList) PrintList() []int {
	array := make([]int, 0, l.Length)
	current := l.Head
	for current != nil {
		array = append(array, current.Value)
		current = current.Next
	}
	return array
}

func (l *LinkedList) TraverseToIndex(index int) *Node {
	if index < 0 {
		return nil
	}
	counter := 0
	current := l.Head
	for current != nil && counter != index {
		current = current.Next
		counter++
	}
	return current
}

func (l *LinkedList) Insert(index int, value int) []int {
	if index <= 0 {
		l.Prepend(value)
		return l.PrintList()
	}
	if index >= l.Length {
		l.Append(value)
		return l.PrintList()
	}

	newNode := &Node{Value: value}
	leader := l.TraverseToIndex(index - 1)
	if leader == nil {
		// fallback: append
		l.Append(value)
		return l.PrintList()
	}
	holdingPointer := leader.Next
	leader.Next = newNode
	newNode.Next = holdingPointer
	l.Length++
	return l.PrintList()
}

func (l *LinkedList) Remove(index int) []int {
	if index < 0 || index >= l.Length || l.Length == 0 {
		return l.PrintList()
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Length--
		if l.Length == 0 {
			l.Tail = nil
		}
		return l.PrintList()
	}

	leader := l.TraverseToIndex(index - 1)
	if leader == nil || leader.Next == nil {
		return l.PrintList()
	}
	unwantedNode := leader.Next
	leader.Next = unwantedNode.Next
	if index == l.Length-1 {
		l.Tail = leader
	}
	l.Length--
	return l.PrintList()
}

func main() {
	myLinkedList := NewLinkedList(10)
	myLinkedList.Append(5)
	myLinkedList.Append(16)
	myLinkedList.Prepend(1)
	fmt.Println("After appends/prepend:", myLinkedList.PrintList())

	myLinkedList.Insert(2, 99)
	fmt.Println("After insert(2,99):", myLinkedList.PrintList())

	myLinkedList.Insert(20, 88)
	fmt.Println("After insert(20,88):", myLinkedList.PrintList())

	myLinkedList.Remove(2)
	fmt.Println("After remove(2):", myLinkedList.PrintList())
}
