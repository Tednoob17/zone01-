package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			fprime(i)
		}
	}
}

func fprime(i int) {
	if i == 1 {
		return
	}
	div := 2
	for i > 1 {
		if i%div == 0 {
			fmt.Print(div)
			i = i / div

			if i > div {
				fmt.Print("*")
			}
			div--
		}
		div++
	}
	fmt.Println()
}
