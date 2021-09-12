package main

import "fmt"

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	result := nb - 1
	for i := result; i >= 1; i-- {
		nb *= i
	}
	return nb
}
