package main

import "fmt"

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}

func IsPrintable(s string) bool {
	for _, charAscii := range s {
		if (charAscii < 'a' || charAscii > 'z') && (charAscii < 'A' || charAscii > 'Z') && (charAscii < '0' || charAscii > '9') {
			return false
		}
	}
	return true
}
