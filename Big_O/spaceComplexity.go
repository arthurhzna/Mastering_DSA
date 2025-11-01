package main

import "fmt"

// #5 Space complexity O(1)
// boooo prints "booooo" n times.
func boooo(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("booooo")
	}
}

// #6 Space complexity O(n)
func arrayOfHiNTimes(n int) []string {
	hiArray := make([]string, n)
	for i := 0; i < n; i++ {
		hiArray[i] = "hi"
	}
	return hiArray
}

func main() {
	fmt.Println("Space complexity O(1)")
	boooo(6)
	fmt.Println("Space complexity O(n)")
	fmt.Println(arrayOfHiNTimes(6))
}
