package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintComb2()
}

func PrintComb2() {
	for i := '0'; i <= 98; i++ {
		for b := i + 1; b <= 99; b++ {
			z01.PrintRune(rune(i/10 + '0'))
			z01.PrintRune(rune(i%10 + '0'))
			z01.PrintRune(' ')
			z01.PrintRune(rune(b/10 + '0'))
			z01.PrintRune(rune(b/10 + '0'))
			if a != 98 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
