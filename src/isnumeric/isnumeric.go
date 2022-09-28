package main

import "fmt"

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}

func IsNumeric(s string) bool {
	for _, charAscii := range s {
		if charAscii <= '0' || charAscii >= '9' {
			return false
		}
	}
	return true
}
