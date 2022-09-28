package main

import "fmt"

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	i := 0
	for ; i < len(s); i++ {
	}
	return i
}
