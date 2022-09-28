package main

import "fmt"

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

func AppendRange(min, max int) []int {
	var result []int
	if min > max {
		return result
	}
	for elem := min - 1; elem < max-1; elem++ {
		result = append(result, elem+1)
	}
	return result
}
