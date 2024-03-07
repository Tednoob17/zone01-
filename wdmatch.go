package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	if wdmatch(s2, s2) {
		fmt.Println(s1)
	}
}

func wdmatch(s1, s2 string) bool {
	j := 0
	for _, char := range s1 {
		found := false
		for ; j < len(s2); j++ {
			if rune(s2[j]) == char {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Check if the number of command-line arguments is exactly 3
// 	if len(os.Args) != 3 {
// 		return // Exit if the number of arguments is not 3
// 	}

// 	// Retrieve the two input strings from command-line arguments
// 	s1 := os.Args[1]
// 	s2 := os.Args[2]

// 	// Check if it's possible to write the first string with characters from the second string
// 	if matchWords(s1, s2) {
// 		fmt.Println(s1) // If so, print the first string
// 	}
// }

// // matchWords checks whether it's possible to write the first string with characters from the second string
// func matchWords(s1, s2 string) bool {
// 	s2Index := 0 // Initialize the index for the second string to 0

// 	// Loop over each character in the first string
// 	for _, c1 := range s1 {
// 		found := false // Initialize a flag to track if the character is found in the second string

// 		// Iterate through the second string starting from the current index
// 		for ; s2Index < len(s2); s2Index++ {
// 			// If the current character in the second string matches the character from the first string
// 			if rune(s2[s2Index]) == c1 {
// 				found = true // Set the flag to true, indicating the character is found
// 				break        // Exit the loop, as we found the character in the second string
// 			}
// 		}

// 		// If the character from the first string is not found in the second string
// 		if !found {
// 			return false // Return false, indicating it's not possible to write the first string with characters from the second string
// 		}
// 	}

// 	return true // Return true if all characters from the first string are found in the second string
// }
