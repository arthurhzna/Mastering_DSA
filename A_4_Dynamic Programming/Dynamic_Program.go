package main

import "fmt"

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233...

var calculations = 0

// O(2^n) - Exponential time complexity
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// O(n) - Linear time complexity
func fibonacciMaster() func(int) int {
	cache := make(map[int]int)

	var fib func(int) int
	fib = func(n int) int {
		calculations++
		if val, exists := cache[n]; exists {
			return val
		} else {
			if n < 2 {
				return n
			} else {
				cache[n] = fib(n-1) + fib(n-2)
				return cache[n]
			}
		}
	}

	return fib
}

// O(n) - Linear time complexity
func fibonacciMaster2(n int) int {
	if n < 2 {
		return n
	}

	answer := []int{0, 1}
	for i := 2; i <= n; i++ {
		answer = append(answer, answer[i-2]+answer[i-1])
	}
	return answer[len(answer)-1]
}

func main() {
	fasterFib := fibonacciMaster()

	fmt.Println("Slow", fibonacci(35))
	fmt.Println("DP", fasterFib(100))
	fmt.Println("DP2", fibonacciMaster2(100))
	fmt.Printf("we did %d calculations\n", calculations)
}
