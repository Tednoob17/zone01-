package main

import (
	"fmt"
)

func WordFlip(str string) string {
	str = trimtSpace(str)

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
	var words []string
	start := 0
	for i, char := range str {
		if char == ' ' {
			words = append(words, str[start:i])
			start = i + 1
		}
	}
	words = append(words, str[start:])
	return words
}

func trimtSpace(str string) string {
	start := 0
	end := len(str) - 1

	for start < len(str) && str[start] == ' ' {
		start++
	}
	for end >= start && str[end] == ' ' {
		end--
	}
	return str[start : end+1]
}

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
