package main

import "fmt"

type PriorityQueue struct {
	heap []int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{heap: make([]int, 0)}
}

func (pq *PriorityQueue) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (pq *PriorityQueue) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (pq *PriorityQueue) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (pq *PriorityQueue) hasLeftChild(index int) bool {
	return pq.getLeftChildIndex(index) < len(pq.heap)
}

func (pq *PriorityQueue) hasRightChild(index int) bool {
	return pq.getRightChildIndex(index) < len(pq.heap)
}

func (pq *PriorityQueue) hasParent(index int) bool {
	return pq.getParentIndex(index) >= 0
}

func (pq *PriorityQueue) leftChild(index int) int {
	return pq.heap[pq.getLeftChildIndex(index)]
}

func (pq *PriorityQueue) rightChild(index int) int {
	return pq.heap[pq.getRightChildIndex(index)]
}

func (pq *PriorityQueue) parent(index int) int {
	return pq.heap[pq.getParentIndex(index)]
}

func (pq *PriorityQueue) swap(indexOne, indexTwo int) {
	pq.heap[indexOne], pq.heap[indexTwo] = pq.heap[indexTwo], pq.heap[indexOne]
}

func (pq *PriorityQueue) Peek() *int {
	if len(pq.heap) == 0 {
		return nil
	}
	return &pq.heap[0]
}

func (pq *PriorityQueue) Remove() *int {
	if len(pq.heap) == 0 {
		return nil
	}
	item := pq.heap[0]
	pq.heap[0] = pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[:len(pq.heap)-1]
	pq.heapifyDown()
	return &item
}

func (pq *PriorityQueue) Add(item int) {
	pq.heap = append(pq.heap, item)
	pq.heapifyUp()
}

func (pq *PriorityQueue) heapifyUp() {
	index := len(pq.heap) - 1
	for pq.hasParent(index) && pq.parent(index) < pq.heap[index] {
		pq.swap(pq.getParentIndex(index), index)
		index = pq.getParentIndex(index)
	}
}

func (pq *PriorityQueue) heapifyDown() {
	index := 0
	for pq.hasLeftChild(index) {
		smallerChildIndex := pq.getLeftChildIndex(index)
		if pq.hasRightChild(index) && pq.rightChild(index) > pq.leftChild(index) {
			smallerChildIndex = pq.getRightChildIndex(index)
		}
		if pq.heap[index] > pq.heap[smallerChildIndex] {
			break
		} else {
			pq.swap(index, smallerChildIndex)
		}
		index = smallerChildIndex
	}
}

func main() {
	priQueue := NewPriorityQueue()

	priQueue.Add(32)
	priQueue.Add(45)
	priQueue.Add(12)
	priQueue.Add(65)
	priQueue.Add(85)

	fmt.Println(*priQueue.Peek())
	fmt.Println(*priQueue.Remove())
	fmt.Println(*priQueue.Peek())
	fmt.Println(*priQueue.Remove())
	fmt.Println(*priQueue.Peek())
	fmt.Println(*priQueue.Remove())
}
