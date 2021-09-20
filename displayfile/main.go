package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	length := 0
	for i := range args {
		length = i
	}

	if length >= 1 {
		fmt.Println("Too many arguments")
	} else if length == 0 {
		fmt.Println("File name missing")
	} else if args[0] == "quest8.txt" {

		contents, e := ioutil.ReadFile(args[0])
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println(string(contents))
	}
}
