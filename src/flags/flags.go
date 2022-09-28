package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	insert := false
	if len(arguments) > 0 && arguments[0] == "--insert" || arguments[0] == "-i" {
		insert = true
		arguments = arguments[1:]
	}
	for _, arg := range arguments {
		if arg[0] == '-' {
			z01.PrintRune(' ')
			continue
		}
		if insert {
			z01.PrintRune(' ')
		}
		for _, char := range arg {
			z01.PrintRune(char)
		}
		nbr := 0
		for _, letter := range arg {
			if letter < '0' || letter > '9' {
				z01.PrintRune(' ')
				nbr = 0
				break
			}
			nbr = nbr*10 + int(letter-'0')
		}
		if nbr > 0 {
			z01.PrintRune(rune(nbr + 'a' - 1))
		}
	}
}
