package main

import (
	"go/printer"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	if n < 0{
		z01.PrintRune('-')
		Printnbr(-n)

	}else if n > 9 {
		PrintNbr(n / 10)
		printNbr(n % 10)
	}else {
		z01.PrintRune(rune( n + '0'))
	}

}