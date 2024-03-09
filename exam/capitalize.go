package main

import "fmt"

func Capitalize(s string) string {
	a := []rune(s)
	l := 0
	for i := range a {
		if (a[i] >= 'a' && a[i] <= 'z') || (a[i] >= 'A' && a[i] <= 'Z') || (a[i] <= '9' && a[i] >= '0') {
			l++
		} else {
			l = 0
		}

		if l == 1 && a[i] >= 'a' && a[i] <= 'z' {
			a[i] -= 32

		} else if l > 1 && a[i] >= 'A' && a[i] <= 'Z' {
			a[i] += 32
		}
	}
	s = string(a)
	return s
}
func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
