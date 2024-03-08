package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))

	fmt.Println(Itoa(987654321))
}
func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sine := ""
	if n < 0 {
		sine = "-"
		n *= -1
	}
	result := ""
	for n > 0 {
		digit := n % 10
		result = string(digit+'0') + result
		n = n / 10
	}

	return sine + result
}
