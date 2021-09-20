package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]

	length := 0
	for i := range arguments {
		length = i + 1
	}

	if length > 1 {
		fmt.Println("Too many arguments")
	} else if length == 0 {
		fmt.Println("File name missing")
	} else if arguments[0] == "quest8.txt" {

		contents, e := ioutil.ReadFile(arguments[0])
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println(string(contents))

	}
}
