package main

import "fmt"

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}

func ToUpper(s string) string {
	var str = []rune(s)
	for i, charAscii := range s {
		if charAscii >= 65 && charAscii <= 90 {
			str[i] += 32
		}
	}
	return string(str)
}
