package main

import "fmt"

func PrintComb2() {
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			fmt.Printf("%02d %02d", i, j)
			if !(i == 98 && j == 99) {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()
}

func main() {
	PrintComb2()
}
