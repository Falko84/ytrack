package main

import "fmt"

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

func Index(s string, toFind string) int {
	var Compteur = []rune(toFind)
	for i := 0; i != len(s); i++ {
		if rune(s[i]) == Compteur[0] {
			if len(toFind) == 1 {
				return i
			}
			for j := 0; j != len(toFind)-1; {
				if rune(s[i+j]) == Compteur[j] && rune(s[i+j+1]) == Compteur[j+1] && len(toFind) != 1 && len(toFind)-1 == j+1 {
					return i
				}
				if rune(s[i+j]) == Compteur[j] {
					j++
				} else {
					break
				}
			}
		}
	}
	return -1
}
