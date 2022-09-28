package main

import "fmt"

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	var res int
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			res = i
			break
		}
	}
	return res
}
