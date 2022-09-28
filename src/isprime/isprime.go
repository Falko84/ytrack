package main

import "fmt"

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}

func IsPrime(nb int) bool {
	if nb == 0 {
		return false
	}
	if nb == 1 {
		return false
	}

	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
