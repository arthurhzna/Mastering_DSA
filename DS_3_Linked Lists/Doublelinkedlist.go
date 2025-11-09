package main

import "fmt"

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewDoublyLinkedList(value int) *DoublyLinkedList {
	node := &Node{Value: value}
	return &DoublyLinkedList{Head: node, Tail: node, Length: 1}
}

func (l *DoublyLinkedList) Append(value int) *DoublyLinkedList {
	newNode := &Node{Value: value, Prev: l.Tail}
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

func (l *DoublyLinkedList) Prepend(value int) *DoublyLinkedList {
	newNode := &Node{Value: value, Next: l.Head}
	if l.Length == 0 {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Head.Prev = newNode
		l.Head = newNode
	}
	l.Length++
	return l
}

func (l *DoublyLinkedList) PrintList() []int {
	out := make([]int, 0, l.Length)
	for cur := l.Head; cur != nil; cur = cur.Next {
		out = append(out, cur.Value)
	}
	return out
}

func (l *DoublyLinkedList) TraverseToIndex(index int) *Node {
	if index < 0 || index >= l.Length {
		return nil
	}
	// choose direction based on index for efficiency
	if index < l.Length/2 {
		counter := 0
		cur := l.Head
		for cur != nil && counter != index {
			cur = cur.Next
			counter++
		}
		return cur
	}
	counter := l.Length - 1
	cur := l.Tail
	for cur != nil && counter != index {
		cur = cur.Prev
		counter--
	}
	return cur
}

func (l *DoublyLinkedList) Insert(index int, value int) []int {
	if index <= 0 {
		l.Prepend(value)
		return l.PrintList()
	}
	if index >= l.Length {
		l.Append(value)
		return l.PrintList()
	}
	leader := l.TraverseToIndex(index - 1)
	if leader == nil || leader.Next == nil {
		l.Append(value)
		return l.PrintList()
	}
	newNode := &Node{Value: value}
	follower := leader.Next

	leader.Next = newNode
	newNode.Prev = leader

	newNode.Next = follower
	follower.Prev = newNode

	l.Length++
	return l.PrintList()
}

func (l *DoublyLinkedList) Remove(index int) []int {
	if l.Length == 0 || index < 0 || index >= l.Length {
		return l.PrintList()
	}
	if index == 0 {
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		} else {
			l.Tail = nil
		}
		l.Length--
		return l.PrintList()
	}
	if index == l.Length-1 {
		l.Tail = l.Tail.Prev
		if l.Tail != nil {
			l.Tail.Next = nil
		} else {
			l.Head = nil
		}
		l.Length--
		return l.PrintList()
	}
	unwanted := l.TraverseToIndex(index)
	if unwanted == nil {
		return l.PrintList()
	}
	prev := unwanted.Prev
	next := unwanted.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
	l.Length--
	return l.PrintList()
}

func main() {
	l := NewDoublyLinkedList(10)
	l.Append(5)
	l.Append(16)
	l.Prepend(1)
	fmt.Println("After append/prepend:", l.PrintList())

	l.Insert(2, 99)
	fmt.Println("After insert(2,99):", l.PrintList())

	l.Insert(20, 88)
	fmt.Println("After insert(20,88):", l.PrintList())

	l.Remove(2)
	fmt.Println("After remove(2):", l.PrintList())

	l.Remove(0)
	fmt.Println("After remove(0):", l.PrintList())

	l.Remove(l.Length - 1)
	fmt.Println("After remove(tail):", l.PrintList())
}
