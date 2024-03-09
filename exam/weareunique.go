package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
func WeAreUnique(str1, str2 string) int {
	charCount := make(map[rune]int)
	for _, i := range str1 {
		charCount[i]++
	}
	for _, char := range str2 {
		charCount[char]++
	}
	unque := 0
	for _, i := range charCount {
		if i == 1 {
			unque++
		}
	}
	return unque

}
