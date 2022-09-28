package main

import "fmt"

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}

func IsAlpha(s string) bool {
	for _, charAscii := range s {
		if (charAscii < 'a' || charAscii > 'z') && (charAscii < 'A' || charAscii > 'Z') && (charAscii < '0' || charAscii > '9') {
			return false
		}
	}
	return true
}
