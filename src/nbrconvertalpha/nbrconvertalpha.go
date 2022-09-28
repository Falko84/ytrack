package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	upper := false
	if len(arguments) > 0 && arguments[0] == "--upper" {
		upper = true
		arguments = arguments[1:]
	}
	for _, arg := range arguments {
		if arg[0] == '-' {
			z01.PrintRune(' ')
			continue
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
			if upper {
				z01.PrintRune(rune(nbr + 'A' - 1))
			} else {
				z01.PrintRune(rune(nbr + 'a' - 1))
			}
		}
	}
	z01.PrintRune(10)
}
