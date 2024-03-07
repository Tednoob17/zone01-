package main

import (
	"fmt"
)

func trimSpaces(str string) string {
	start, end := 0, len(str)-1

	// Find start index of non-space characters
	for start < len(str) && str[start] == ' ' {
		start++
	}

	// Find end index of non-space characters
	for end >= 0 && str[end] == ' ' {
		end--
	}

	// Return substring with trimmed spaces
	return str[start : end+1]
}

func main() {
	str := "   Hello, world!   "
	trimmed := trimSpaces(str)
	fmt.Println(trimmed) // Output: "Hello, world!"
}
