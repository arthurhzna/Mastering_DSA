package main

import "fmt"

func fibonacciIterative(n int) int {
	arr := []int{0, 1}
	for i := 2; i < n+1; i++ {
		arr = append(arr, arr[i-2]+arr[i-1])
	}
	return arr[n]
}

func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func main() {
	fmt.Println(fibonacciIterative(3))
	fmt.Println(fibonacciRecursive(6))
}
