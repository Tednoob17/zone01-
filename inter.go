package main

import (
	"fmt"
	"os"
)

func containsRune(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}

func commonCharacters(str1, str2 string) string {
	result := ""
	for _, char := range str1 {
		if containsRune(str2, char) && !containsRune(result, char) {
			result += string(char)
		}
	}
	return result
}

func main() {
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		fmt.Println(commonCharacters(str1, str2))
	}
}
