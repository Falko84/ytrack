package main

import "fmt"

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func IsNumeric(s string) bool {
	for _, charAscii := range s {
		if charAscii <= '0' || charAscii >= '9' {
			return false
		}
	}
	return true
}

func CountIf(f func(string) bool, tab []string) int {
	var count int = 0
	for _, i := range tab {
		if f(i) {
			count += 1
		}
	}
	return count
}
