package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {
	var res int
	for _, v := range s {
		res = res*10 + int(v-'0')
	}
	return res
}
