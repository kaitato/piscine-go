package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	var i, j int

	for i = 10; i <= 10; i++ {
		for j = 1; j <= 10; j++ {
			fmt.Println(i, " * ", j, " = ", i*j)
		}
	}
}
