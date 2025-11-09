package main

import "fmt"

func quickSort(array []int, left int, right int) []int {
	var pivot int
	var partitionIndex int

	if left < right {
		pivot = right
		partitionIndex = partition(array, pivot, left, right)

		quickSort(array, left, partitionIndex-1)
		quickSort(array, partitionIndex+1, right)
	}
	return array
}

func partition(array []int, pivot int, left int, right int) int {
	pivotValue := array[pivot]
	partitionIndex := left

	for i := left; i < right; i++ {
		if array[i] < pivotValue {
			swap(array, i, partitionIndex)
			partitionIndex++
		}
	}
	swap(array, right, partitionIndex)
	return partitionIndex
}

func swap(array []int, firstIndex int, secondIndex int) {
	temp := array[firstIndex]
	array[firstIndex] = array[secondIndex]
	array[secondIndex] = temp
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
}
