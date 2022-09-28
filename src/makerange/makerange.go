package main

import "fmt"

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

func AppendRange(min, max int) []int {
	var nil []int
	if min > max {
		return nil
	}
	result := make([]int, max-min)
	var temp int = min
	for elem := 0; elem != max-min; elem++ {
		result[elem] = temp
		temp++
	}
	return result
}
