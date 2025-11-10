package main

import "fmt"

func reverseString(str string) string {
	arrayStr := []rune(str)
	reversedArray := []rune{}

	var addToArray func(array []rune)
	addToArray = func(array []rune) {
		if len(array) > 0 {
			lastIndex := len(array) - 1
			reversedArray = append(reversedArray, array[lastIndex])
			addToArray(array[:lastIndex])
		}
		return
	}

	addToArray(arrayStr)
	return string(reversedArray)
}

func reverseStringRecursive(str string) string {
	if str == "" {
		return ""
	}
	return reverseStringRecursive(str[1:]) + string(str[0])
}

func main() {
	result1 := reverseString("yoyo master")
	fmt.Println("reverseString:", result1)

	result2 := reverseStringRecursive("yoyo master")
	fmt.Println("reverseStringRecursive:", result2)
}
