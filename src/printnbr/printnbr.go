package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var i int
	var n2 int = n
	var modu int = 1
	if n < 0 {
		n2 *= -1
		n *= -1
		z01.PrintRune('-')
	}
	for i = 0; n2 > 0; i++ {
		n2 /= 10
		modu *= 10
	}
	println(modu)
	for j := 0; j != i; j++ {

	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
