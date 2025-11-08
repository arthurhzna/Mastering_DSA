package main

import "fmt"

type CrazyQueue struct {
	first []interface{}
	last  []interface{}
}

func NewCrazyQueue() *CrazyQueue {
	return &CrazyQueue{
		first: make([]interface{}, 0),
		last:  make([]interface{}, 0),
	}
}

func (cq *CrazyQueue) Enqueue(value interface{}) *CrazyQueue {
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

func (cq *CrazyQueue) Dequeue() *CrazyQueue {
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

func (cq *CrazyQueue) Peek() interface{} {
	if len(cq.first) > 0 {
		return cq.first[len(cq.first)-1]
	}
	if len(cq.last) > 0 {
		return cq.last[0]
	}
	return nil
}

func main() {
	myQueue := NewCrazyQueue()

	fmt.Println(myQueue.Peek())

	myQueue.Enqueue("Joy")
	myQueue.Enqueue("Matt")
	myQueue.Enqueue("Pavel")

	fmt.Println(myQueue.Peek())
	fmt.Println("========")

	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Dequeue())

	fmt.Println("========")
	fmt.Println(myQueue.Peek())
	fmt.Printf("%+v\n", myQueue)
}
