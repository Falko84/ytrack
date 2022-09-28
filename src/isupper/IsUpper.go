package main

import "fmt"

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}

func IsUpper(s string) bool {
	for _, charAscii := range s {
		if charAscii <= 'A' || charAscii >= 'Z' {
			return false
		}
	}
	return true
}
