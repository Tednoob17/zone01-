package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
func ConcatSlice(slice1, slice2 []int) []int {
	var result []int
	for _, i := range slice1 {
		result = append(result, i)
	}
	for _, i := range slice2 {
		result = append(result, i)
	}

	return result
}
