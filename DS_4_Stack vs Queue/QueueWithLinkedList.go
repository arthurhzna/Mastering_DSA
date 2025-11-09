package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

func (q *Queue) Peek() *Node {
	return q.first
}

func (q *Queue) Enqueue(value interface{}) *Queue {
	newNode := NewNode(value)
	if q.length == 0 {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++
	return q
}

func (q *Queue) Dequeue() *Queue {
	if q.first == nil {
		return nil
	}
	if q.first == q.last {
		q.last = nil
	}
	q.first = q.first.next
	q.length--
	return q
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) String() string {
	if q.first == nil {
		return "Queue is empty"
	}

	result := "Queue: "
	current := q.first
	for current != nil {
		result += fmt.Sprintf("%v", current.value)
		if current.next != nil {
			result += " -> "
		}
		current = current.next
	}
	return result
}

func main() {
	myQueue := NewQueue()
	myQueue.Enqueue("Joy")
	myQueue.Enqueue("Matt")
	myQueue.Enqueue("Pavel")
	myQueue.Enqueue("Samir")
	myQueue.Dequeue()

	fmt.Println("Peek:", myQueue.Peek().value)
	fmt.Println(myQueue)
}
