package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	res := ""
	for j := 0; j < len(arr); j++ {
		res = converthex(int(arr[j]))
		printstr(res)
		if ((j+1)%4 == 0 && j != 0) || j == len(arr)-1 {
			z01.PrintRune('\n')

		} else {
			z01.PrintRune(' ')
		}

	}
	for _, v := range arr {
		if v >= 33 && v <= 126 {
			printstr(string(v))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')

}

func converthex(nb int) string {
	tab := ""
	base := "0123456789abcdef"
	if nb == 0 {
		tab = "00" + tab
		return tab
	}
	for nb != 0 {
		tab = string(base[nb%16]) + tab
		nb = nb / 16
	}

	if len(tab) == 1 {
		return "0" + tab
	}

	return tab
}

func printstr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
