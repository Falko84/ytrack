package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

func Map(f func(int) bool, a []int) []bool {
	var ui []bool
	for _, e := range a {
		ui = append(ui, f(e))
	}
	return ui
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
