package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	count := 0
	for s := range arg {
		count = s + 1
	}
	for j := count - 1; j >= 1; j-- {
		for _, a := range arg[j] {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
