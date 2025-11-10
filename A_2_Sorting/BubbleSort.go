package main

import "fmt"

func bubbleSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	bubbleSort(numbers)
	fmt.Println(numbers)
}
