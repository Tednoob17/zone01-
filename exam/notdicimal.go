package main

import "fmt"

func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}

	var num float64
	var isNegative bool
	var err error
	for i, c := range dec {
		if i == 0 && c == '-' {
			isNegative = true
		} else if c < '0' || c > '9' {
			if c != '.' {
				return dec + "\n"
			}
		} else if err == nil {
			num = num*10 + float64(c-'0')
		}
	}

	if !containsDecimal(dec) || dec[len(dec)-1] == '.' || (len(dec) > 1 && dec[len(dec)-2] == '.' && dec[len(dec)-1] == '0') {
		return dec + "\n"
	}

	integerPart := int(num)
	if isNegative {
		integerPart *= -1
	}
	return fmt.Sprintf("%d\n", integerPart)
}

func containsDecimal(dec string) bool {
	for _, c := range dec {
		if c == '.' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
