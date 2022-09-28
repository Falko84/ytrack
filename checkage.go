package main

import "fmt"

func CheckAge(age int) bool {
	// if age >= 18 {
	// 	fmt.Println("tu peux entrer")
	// } else {
	// 	fmt.Println("Sortez")
	// }
	if age >= 18 {
		return true
	} else {
		return false
	}
}

func Checkid(age int, sexe string) string {
	mavariable := ""
	if CheckAge(age) {
		if sexe == "femme" {
			fmt.Println("c'est 10 euro")

		}
		if sexe == "homme" {
			fmt.Println("c'est 15 euro")
		}
	} else {
		mavariable = "sortez"
	}
	return mavariable
}

func main() {
	fmt.Println(Checkid(19, "femme"))
	fmt.Println(Checkid(9, "homme"))
	fmt.Println(Checkid(25, "homme"))
}
