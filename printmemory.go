package piscine

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintAdress(c rune) {
	toint := int(c)
	if toint == 0 {
		z01.PrintRune('0')
	}
	PrintHex(toint)
}

func PrintHex(nbr int) {
	base := "0123456789abcdef"
	if nbr >= len(base) {
		PrintHex(nbr / len(base))
		PrintHex(nbr % len(base))
	} else if nbr <= len(base)-1 && nbr >= 0 {
		z01.PrintRune(rune(base[nbr]))
	}
}

func PrintMemory(arr [10]byte) {
	i := 0
	for i < len(arr) {
		for j := 0; j < 4 && i < len(arr); j++ {
			runes := rune(arr[i])
			PrintAdress(runes)
			if j < 3 && i < len(arr)-1 {
				z01.PrintRune(' ')
			}
			i++
		}
		z01.PrintRune('\n')
	}
	for i := 0; i < len(arr); i++ {
		if unicode.IsGraphic(rune(arr[i])) {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
