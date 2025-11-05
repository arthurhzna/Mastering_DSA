package main

import (
	"fmt"
)

type MyArray struct { // like class
	length int           // like properties in class
	data   []interface{} // like properties in class
}

func NewMyArray() *MyArray { // initialize a new array like class constructor
	return &MyArray{
		length: 0,
		data:   make([]interface{}, 0),
	}
}

func (a *MyArray) Get(index int) interface{} { // like method in class
	if index < 0 || index >= a.length { // a like this is the object of the class
		return nil
	}
	return a.data[index]
}

func (a *MyArray) Push(item interface{}) int {
	a.data = append(a.data, item)
	a.length++
	return a.length
}

func (a *MyArray) Pop() interface{} {
	if a.length == 0 {
		return nil
	}
	last := a.data[a.length-1]
	a.data = a.data[:a.length-1]
	a.length--
	return last
}

func (a *MyArray) Delete(index int) interface{} {
	if index < 0 || index >= a.length {
		return nil
	}
	item := a.data[index]
	a.shiftItems(index)
	return item
}

func (a *MyArray) shiftItems(index int) {
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.data = a.data[:a.length-1]
	a.length--
}

func main() {
	newArray := NewMyArray() // like instance of class

	newArray.Push("hi")
	newArray.Push("you")
	newArray.Push("!")
	fmt.Println(newArray.Get(0))

	newArray.Pop()
	newArray.Delete(1)

	fmt.Printf("length=%d data=%v\n", newArray.length, newArray.data)
}
