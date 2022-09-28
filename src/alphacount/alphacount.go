package main

import (
	"fmt"
)

func main() {
	s := "Hello 78 World!	4455/"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

func AlphaCount(s string) int {
	var compteur int
	for _, charAscii := range s {
		if (charAscii <= 'a' || charAscii >= 'z') && (charAscii <= 'A' || charAscii >= 'Z') {
			compteur++
		}
	}
	return compteur - 1
}
