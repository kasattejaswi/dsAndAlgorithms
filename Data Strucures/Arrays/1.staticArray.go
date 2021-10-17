package main

import "fmt"

// Creating a new array with values already present
var staticArr = [5]int{32, 12, 45, 3, 12}

// Creating an empty array
var staticEmptArr [5]int

func main() {
	// Printing static array with values already filled
	fmt.Println(staticArr)

	// Printing static empty array
	fmt.Println(staticEmptArr)

	// Adding value at a particular index
	staticEmptArr[1] = 22
	fmt.Println(staticEmptArr)
}
