package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	arguments = tri(arguments)
	for _, arg := range arguments {
		for _, letter := range arg {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}
}

func tri(args []string) []string {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	return args
}
