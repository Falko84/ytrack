package main

import "fmt"

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

func ConcatParams(args []string) string {
	var result string
	for _, str := range args {
		result += str + "\n"
	}
	return result
}
