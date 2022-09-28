package main

import "fmt"

func main() {
	a := 4
	maFonction(&a)
	fmt.Println(a)
}

func maFonction(b *int) *int {
	*b *= *b
	return b
}
