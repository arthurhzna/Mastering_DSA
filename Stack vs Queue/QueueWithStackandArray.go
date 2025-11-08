package main

import "fmt"

type CrazyQueue[T any] struct {
	first []T
	last  []T
}

func NewCrazyQueue[T any]() *CrazyQueue[T] {
	return &CrazyQueue[T]{
		first: make([]T, 0),
		last:  make([]T, 0),
	}
}

func (cq *CrazyQueue[T]) Enqueue(value T) *CrazyQueue[T] {
	length := len(cq.first)
	
	for i := 0; i < length; i++ {
		if len(cq.first) > 0 {
			lastIndex := len(cq.first) - 1
			element := cq.first[lastIndex]
			cq.first = cq.first[:lastIndex]
			cq.last = append(cq.last, element)
		}
	}
	
	cq.last = append(cq.last, value)
	return cq
}

func (cq *CrazyQueue[T]) Dequeue() *CrazyQueue[T] {
	length := len(cq.last)
	
	for i := 0; i < length; i++ {
		if len(cq.last) > 0 {
			lastIndex := len(cq.last) - 1
			element := cq.last[lastIndex]
			cq.last = cq.last[:lastIndex]
			cq.first = append(cq.first, element)
		}
	}
	
	if len(cq.first) > 0 {
		cq.first = cq.first[:len(cq.first)-1]
	}
	
	return cq
}

func (cq *CrazyQueue[T]) Peek() *T {
	if len(cq.first) > 0 {
		return &cq.first[len(cq.first)-1]
	}
	if len(cq.last) > 0 {
		return &cq.last[0]
	}
	return nil
}

func main() {
	myQueue := NewCrazyQueue[string]()
	
	if peek := myQueue.Peek(); peek != nil {
		fmt.Println(*peek)
	} else {
		fmt.Println("<nil>")
	}
	
	myQueue.Enqueue("Joy")
	myQueue.Enqueue("Matt")
	myQueue.Enqueue("Pavel")
	
	if peek := myQueue.Peek(); peek != nil {
		fmt.Println(*peek)
	}
	
	fmt.Println("========")
	fmt.Printf("%+v\n", myQueue.Dequeue())
	fmt.Printf("%+v\n", myQueue.Dequeue())
	fmt.Printf("%+v\n", myQueue.Dequeue())
	fmt.Println("========")
	
	if peek := myQueue.Peek(); peek != nil {
		fmt.Println(*peek)
	} else {
		fmt.Println("<nil>")
	}
	
	fmt.Printf("%+v\n", myQueue)
}