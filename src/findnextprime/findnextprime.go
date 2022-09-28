package main

import "fmt"

func main() {
	fmt.Println(FindNextPrime(73))
	fmt.Println(FindNextPrime(777))
}

func IsPrime(nb int) bool {
	if nb == 0 {
		return false
	}

	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		nb++
	}
	for ; IsPrime(nb) != true; nb++ {
	}
	return nb
}
