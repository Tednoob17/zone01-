package main

import (
	"fmt"
)

func trimSpaces(str string) string {
	start, end := 0, len(str)-1

	for start < len(str) && str[start] == ' ' {
		start++
	}

	for end >= 0 && str[end] == ' ' {
		end--
	}

	return str[start : end+1]
}

func main() {
	str := "   Hello, world!   "
	trimmed := trimSpaces(str)
	fmt.Println(trimmed)
}
