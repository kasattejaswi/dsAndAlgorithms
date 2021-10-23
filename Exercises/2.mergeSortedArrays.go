package main

import "fmt"

func main() {
	arr1 := [5]int{2, 5, 7, 18, 20}
	arr2 := [6]int{1, 3, 15, 18, 20, 23}
	fmt.Println(mergeSortedArrays(arr1[:], arr2[:]))
}

// Merge two sorted arrays
func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	mergedArray := []int{}
	arr1Index := 0
	arr2Index := 0
	for true {
		if arr1Index > len(arr1)-1 && arr2Index > len(arr2)-1 {
			break
		} else if arr1Index > len(arr1)-1 {
			mergedArray = append(mergedArray, arr2[arr2Index])
			arr2Index++
		} else if arr2Index > len(arr2)-1 {
			mergedArray = append(mergedArray, arr1[arr1Index])
			arr1Index++
		} else if arr1[arr1Index] < arr2[arr2Index] {
			mergedArray = append(mergedArray, arr1[arr1Index])
			arr1Index++
		} else {
			mergedArray = append(mergedArray, arr2[arr2Index])
			arr2Index++
		}
	}
	return mergedArray
}
