package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {

	if len(str) == 0 {
		return "invalid"
	}
	str = trimSpace(str)
	words := splitwithSpace(str)
	reverseWords(words)
	revers := joinWords(words)

	return revers + "\n"
}
func joinWords(words []string) string {
	var result string

	for i, word := range words {
		result += word
		if i != len(words)-1 {
			result += " "
		}
	}
	return result
}

func reverseWords(words []string) {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
}
func splitwithSpace(str string) []string {
	var word []string
	start := 0
	for i, char := range str {
		if char == ' ' {
			word = append(word, str[start:i])
			start = 1 + i
		}
	}
	word = append(word, str[start:])
	return word
}
func trimSpace(str string) string {
	start, end := 0, len(str)-1
	for start < len(str) && str[start] == ' ' {
		start++
	}
	for end >= start && str[end] == ' ' {
		end--
	}
	return str[start : end+1]
}
