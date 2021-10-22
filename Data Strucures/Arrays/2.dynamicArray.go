package main

import (
	"errors"
	"fmt"
)

type DArray struct {
	currentLen int
	// TODO Replace map with bucket linked list
	currentArr map[int]int
}

func main() {
	myDArr := DArray{}
	myDArr.add(2)
	myDArr.add(34)
	myDArr.add(29)
	myDArr.add(364)
	myDArr.add(21)
	myDArr.add(343)
	fmt.Println("Size of current dynamic array: ", myDArr.size())
	myDArr.print()
	myDArr.pop()
	fmt.Println("Size of current dynamic array: ", myDArr.size())
	myDArr.print()
	err1 := myDArr.delete(0)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("Size of current dynamic array: ", myDArr.size())
	myDArr.print()
	err2 := myDArr.delete(30)
	if err2 != nil {
		fmt.Println(err2)
	}

}

// Add a new element
func (a *DArray) add(n int) {
	if a.currentLen == 0 {
		a.currentArr = make(map[int]int)
	}
	a.currentArr[a.currentLen] = n
	a.currentLen++
}

// Returns the size of dynamic array
func (a *DArray) size() int {
	return a.currentLen
}

// Deletes last element in array
func (a *DArray) pop() {
	delete(a.currentArr, a.currentLen-1)
	a.currentLen--
}

// Deletes element at a particular index
func (a *DArray) delete(index int) error {
	if index > a.currentLen {
		return errors.New("ERROR: Index cannot be greater than the max size")
	}
	delete(a.currentArr, index)
	newmap := make(map[int]int)
	tempIndex := 0
	for i := 0; i < a.currentLen; i++ {
		if val, ok := a.currentArr[i]; ok {
			newmap[tempIndex] = val
			tempIndex++
		}
	}
	a.currentLen--
	a.currentArr = newmap
	return nil
}

// Prints the entire array
func (a *DArray) print() {
	for i := 0; i < a.currentLen; i++ {
		fmt.Print(a.currentArr[i], " ")
	}
	fmt.Println()
}
