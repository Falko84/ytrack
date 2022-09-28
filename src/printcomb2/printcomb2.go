package main

import "github.com/01-edu/z01"

func main() {
	var i int
	var j int
	var nigga bool = true
	for j = 0; j <= 99; j++ {
		for i = j + 1; i <= 99; i++ {
			if !nigga {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			if j < 10 {
				z01.PrintRune(48)
				z01.PrintRune(rune(j + 48))
			} else {
				z01.PrintRune(rune(j/10 + 48))
				z01.PrintRune(rune(j%10 + 48))
			}
			z01.PrintRune(' ')
			if i < 10 {
				z01.PrintRune(48)
				z01.PrintRune(rune(i + 48))
			} else {
				z01.PrintRune(rune(i/10 + 48))
				z01.PrintRune(rune(i%10 + 48))
			}

			nigga = false

		}
	}
}
