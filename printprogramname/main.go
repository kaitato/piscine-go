package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []rune(os.Args[0])
	a := 1
	for i := len(arg) - 1; i >= 0; i-- {
		if arg[i] == '/' {
			a += i
			break
		}
	}
	for j := a; j < len(arg); j++ {
		z01.PrintRune(arg[j])
	}
	z01.PrintRune('\n')
}
