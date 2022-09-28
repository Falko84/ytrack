package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	program := os.Args[0]
	for _, letter := range program {
		z01.PrintRune(letter)
	}
	z01.PrintRune(10)
}
