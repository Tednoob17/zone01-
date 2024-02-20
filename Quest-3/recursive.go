package main

import "fmt"

func main() {
	result := 0 + 1 + 3 + 4 + 5
	fmt.Println(result)
	x := 5
	fmt.Println(x)
}

func reecurx(x int) {
	for i := 0; i <= x; i++ {
		i = i * i
	}
	return
}
