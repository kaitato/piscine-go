package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1]
	result := os.Args[2]

	contents, err := ioutil.ReadFile(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for range contents {
	}
	err = ioutil.WriteFile(result, contents, 0644)
	if err != nil {
		fmt.Println("Error creating", result)
		fmt.Println(err)
		return
	}
}
