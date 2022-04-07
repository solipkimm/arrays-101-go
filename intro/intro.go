package main

import "fmt"

func main() {

	var squareNumber [10]int
	for i := 0; i < 10; i++ {
		var square int = (i + 1) * (i + 1)
		squareNumber[i] = square
	}
	for _, j := range squareNumber {
		fmt.Println(j)
	}

	// Array Capacity
	var dvd [6]string
	var capacity = len(dvd)
	fmt.Println(capacity)

	// Array Length
	var numbers [6]int
	var length int = 0
	// add 3 items into the array
	for i := 0; i < 3; i++ {
		numbers[i] = i * i
		length++
	}
	fmt.Println(length)

}
