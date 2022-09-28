package main

import "fmt"

func main() {
	fmt.Println(RecursivePower(4, 3))
}

func RecursivePower(nb int, power int) int {
	var res int = 1
	if power < 0 {
		return 0
	}

	if power > 0 {
		res = nb * (RecursivePower(nb, power-1))
	}

	return res

}
