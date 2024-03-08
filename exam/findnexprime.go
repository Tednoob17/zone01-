package main

import "fmt"

func FindPrevPrime(nb int) int {
	if nb <= 2 {
		return 0
	}
	if IsPrime(nb) {
		return nb
	}
	nb--
	return FindPrevPrime(nb)

}
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if i%n == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
