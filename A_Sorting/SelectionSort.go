package main

import "fmt"

func selectionSort(array []int) []int {
	length := len(array)
	for i := 0; i < length-1; i++ {
		min := i
		temp := array[i]
		for j := i + 1; j < length; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i] = array[min]
		array[min] = temp
	}
	return array
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(selectionSort(numbers))
}
