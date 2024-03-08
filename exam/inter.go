package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		fmt.Println(findchracter(str1, str2))
	}
}
func findchracter(str1, str2 string) string {
	result := ""
	for _, char := range str1 {
		if continusrune(str2, char) && !continusrune(result, char) {
			result += string(char)
		}
	}
	return result
}
func continusrune(str1 string, char rune) bool {
	for _, i := range str1 {
		if i == char {
			return false
		}
	}
	return true
}
