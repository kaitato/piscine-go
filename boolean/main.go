package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arrayS := []rune(s)
	count := 0
	for range arrayS {
		count++
	}
	for i := 0; i < count; i++ {
		z01.PrintRune(arrayS[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
