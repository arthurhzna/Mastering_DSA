package main

import "fmt"

func insertionSort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		if array[i] < array[0] {
			value := array[i]
			array = append([]int{value}, array[:i]...)
			array = append(array, array[i+1:]...)
		} else {
			for j := 1; j < i; j++ {
				if array[i] > array[j-1] && array[i] < array[j] {
					value := array[i]
					array = append(array[:j], append([]int{value}, array[j:i]...)...)
					array = append(array, array[i+1:]...)
				}
			}
		}
	}
	return array
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(insertionSort(numbers))
}
