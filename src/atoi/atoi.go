package main

import "fmt"

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	var res int
	var sign int = 1
	for i, v := range s {
		if i == 0 && v == '-' {
			sign = -1
			continue
		}
		if i == 0 && v == '+' {
			continue
		}
		if v < '0' || v > '9' {
			return 0
		}
		res = res*10 + int(v-'0')
	}
	return res * sign
}
