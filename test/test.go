package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var result []rune // fmt
	arg := os.Args
	search := []rune(arg[2])
	replace := []rune(arg[3])
	if len(arg) != 4 {
		return
	}
	for _, i := range arg[1] {
		if search[0] == i {
			result = append(result, replace[0]) // fmt
			z01.PrintRune(replace[0])
		} else {
			result = append(result, i) // fmt
			z01.PrintRune(i)
		}
	}
	fmt.Printf(string(result)) // fmt
}
