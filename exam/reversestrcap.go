package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		arg := []rune(arg)
		for i, j := range arg {
			if i+1 == len(arg) || arg[i+1] == ' ' {
				if j >= 'a' && j <= 'z' {
					arg[i] = j - 32
				}
			} else {
				if j >= 'A' && j <= 'Z' {
					arg[i] = j + 32
				}
			}
		}
		fmt.Println(string(arg))
	}
}
