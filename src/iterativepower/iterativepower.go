package main

import "fmt"

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	var temp int = nb
	for ; power != 1; power-- {
		nb = nb * temp
	}
	return nb
}

func main() {
	fmt.Println(IterativePower(4, 3))
}
