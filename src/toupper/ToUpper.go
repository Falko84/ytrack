package main

import "fmt"

func main() {
	fmt.Println(ToUpper("hello! How are you?"))
}

func ToUpper(s string) string {
	var str = []rune(s)
	for i, charAscii := range s {
		if charAscii >= 97 && charAscii <= 122 {
			str[i] -= 32
		}
	}
	return string(str)
}
