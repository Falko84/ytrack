package main

import "fmt"

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}

func IsLower(s string) bool {
	for _, charAscii := range s {
		if charAscii <= 'a' || charAscii >= 'z' {
			return false
		}
	}
	return true
}
