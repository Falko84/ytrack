package main

import "fmt"

func main() {
	fmt.Println(Trimatoi("12345"))
	fmt.Println(Trimatoi("str123ing45"))
	fmt.Println(Trimatoi("012 345"))
	fmt.Println(Trimatoi("Hello World!"))
	fmt.Println(Trimatoi("sd-x1fa2W3s4"))
	fmt.Println(Trimatoi("sdx1-fa2W3s4"))
	fmt.Println(Trimatoi("sdx1+fa2W3s4"))
}

func IsNb(s string) bool {
	var str = []byte(s)
	for i := 0; i != len(s); {
		if str[i] >= 48 && str[i] <= 57 {
			i++
		} else {
			return false
		}
	}
	return true
}

func Trimatoi(s string) int {
	var str = []byte(s)
	var result int = 0
	var isneg bool = false
	for i := 0; i != len(s); i++ {
		if str[i] == 45 && !(result > 0) {
			isneg = true
		}
		if IsNb(string(str[i])) {
			if result != 0 {
				result *= 10
			}
			result += int(str[i] - 48)
		}
	}
	if isneg == true {
		result *= -1
	}
	return result
}
