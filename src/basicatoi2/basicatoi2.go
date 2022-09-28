package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
	var res int
	for _, v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		res = res*10 + int(v-'0')
	}
	return res
}
