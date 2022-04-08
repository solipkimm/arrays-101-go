package main

import "fmt"

func main() {
	// Declare an integer array of 6 elements
	intArray := make([]int, 6)
	length := 0

	// Add 3 elements to the Array
	for i := 0; i < 3; i++ {
		intArray[length] = i
		length++
	}

	printArray(intArray)
	/*
		Index 0 contains 0
		Index 1 contains 1
		Index 2 contains 2
		Index 3 contains 0
		Index 4 contains 0
		Index 5 contains 0
	*/

	// Inserting at the End of an Array

	// It's important to ensure that there is enough space
	//in the array for inserting a new element.
	intArray[length] = 10
	printArray(intArray)
	/*
		Index 0 contains 0
		Index 1 contains 1
		Index 2 contains 2
		Index 3 contains 4
		Index 4 contains 0
		Index 5 contains 0
	*/

	// Inserting at the Start of an Array

	// Insert a new element by shifting each element one index to the right.
	// Go backwards to avoid overwriting any elements.
	for i := 3; i >= 0; i-- {
		intArray[i+1] = intArray[i]
	}

	intArray[0] = 20
	printArray(intArray)
	/*
		Index 0 contains 20
		Index 1 contains 0
		Index 2 contains 1
		Index 3 contains 2
		Index 4 contains 10
		Index 5 contains 0
	*/

	// Inserting Anywhere in the Array

	// Inter the element at index 2
	for i := 4; i >= 2; i-- {
		intArray[i+1] = intArray[i]
	}

	intArray[2] = 30
	printArray(intArray)
	/*
		Index 0 contains 20
		Index 1 contains 0
		Index 2 contains 30
		Index 3 contains 1
		Index 4 contains 2
		Index 5 contains 10
	*/
}

func printArray(intArray []int) {
	for i := 0; i < len(intArray); i++ {
		fmt.Println("Index", i, "contains", intArray[i])
	}
}
