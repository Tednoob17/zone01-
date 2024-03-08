package main

import (
	"fmt"
)

func main() {
	str1 := "zpadinton"
	str2 := "paqefwtdjetyiytjneytjoeyjnejeyj"

	result := findUniqueCharacters(str1, str2)
	fmt.Println(result) // Output: zpadintoqefwjy
}

func main() {
    arg := os.Args[1:]
    if len(arg) != 2 {
        return
    }
    if len(arg) == 2 {
        str1 , str2 := chackingchars(arg[0]), chackingchars(arg[1])
        res := str1 + str2
        res = chackingchars(res)
        fmt.Println(res)
    }
}

func chackingchars(s string) string {
    r := []rune(s)
    for i := 0; i < len(r); i++ {
        for j := i + 1; j < len(r); j++ {
            if r[i] == r[j] {
                r = append(r[:j], r[j+1:]...)
            }
        }
    }
    return string(r)
}