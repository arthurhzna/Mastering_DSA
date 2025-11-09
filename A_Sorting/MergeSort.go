package main

import "fmt"

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	length := len(array)
	middle := length / 2
	left := array[0:middle]
	right := array[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	result := []int{}
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	answer := mergeSort(numbers)
	fmt.Println(answer)
}
