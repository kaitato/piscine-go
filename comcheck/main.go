package main

import (
	"fmt"
	"os"
)

func main() {
	var aRune []rune
	var result string
	word := []string{"01", "galaxy", "galaxy 01"}

	for i := 1; i < len(os.Args); i++ {
		aRune = []rune(os.Args[i])
	}

	for j := 0; j < len(aRune); j++ {
		if aRune[j] != ' ' {
			result += string(aRune[j])
		}
	}

	for _, s := range word {
		if result == s {
			fmt.Println("Alert!!!")
		}
	}
}
