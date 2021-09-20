package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
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

func main() {
	TooMany := "Too many arguments"
	FileMissing := "File name missing"
	arguments := os.Args[1:]

	length := 0
	for i := range arguments {
		length = i + 1
	}

	if length > 1 {
		printstr(TooMany)
	} else if length == 0 {
		printstr(FileMissing)
	} else if arguments[0] == "quest8.txt" {

		content, err := ioutil.ReadFile(arguments[0])
		if err != nil {
			printstr(err.Error())
			return
		}
		printstr(string(content))

	}
}
