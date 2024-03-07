package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintComb2()
}

//	func PrintComb2() {
//		for i := 0; i <= 98; i++ {
//			for b := 1; b <= 99; b++ {
//				z01.PrintRune(rune(i/10 + '0'))
//				z01.PrintRune(rune(i%10 + '0'))
//				z01.PrintRune(' ')
//				z01.PrintRune(rune(b/10 + '0'))
//				z01.PrintRune(rune(b%10 + '0'))
//				if i <= 98 {
//					z01.PrintRune(',')
//					z01.PrintRune(' ')
//				}
//			}
//		}
//		z01.PrintRune('\n')
//	}
func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '8'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					if a < c || (a == c && b < d) {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)

						if !(a == '9' && b == '8' && c == '9' && d == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
}
