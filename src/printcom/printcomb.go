package main

import "github.com/01-edu/z01"

func main() {
	var nigga bool = true
	for i := 48; i != 56; i++ {
		for j := 49; j != 57; j++ {
			for k := 50; k != 58; k++ {
				if i < j && j < k {
					if !nigga {
						z01.PrintRune(rune(','))
						z01.PrintRune(rune(' '))
					}

					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
					nigga = false
				}
			}
		}
	}
}
