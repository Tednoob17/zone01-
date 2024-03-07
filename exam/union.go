package main

import (
	"fmt"
)

func main() {
	str1 := "zpadinton"
	str2 := "paqefwtdjetyiytjneytjoeyjnejeyj"

	result := findUniqueCharacters(str1, str2)
	fmt.Println(result) // Output: zpadintoqefwjy
}

func findUniqueCharacters(str1, str2 string) string {
	characters := make(map[rune]bool)
	order := []rune{}

	// Iterate through first string and add characters to map
	for _, char := range str1 {
		if _, exists := characters[char]; !exists {
			characters[char] = true
			order = append(order, char)
		}
	}

	// Iterate through second string and add characters to map
	for _, char := range str2 {
		if _, exists := characters[char]; !exists {
			characters[char] = true
			order = append(order, char)
		}
	}

	// Form the result string using the order slice
	result := ""
	for _, char := range order {
		result += string(char)
	}

	return result
}
