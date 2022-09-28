package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintParams(args []string) {
	for _, arg := range args {
		for _, letter := range arg {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}
}

func main() {
	PrintParams(os.Args[1:])
}
