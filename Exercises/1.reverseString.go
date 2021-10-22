package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseString("Hi"))
}

func reverseString(s string) string {
	arr := strings.Split(s, "")
	tempStr := ""

	for i := len(arr) - 1; i >= 0; i-- {
		tempStr += arr[i]
	}
	return tempStr
}
