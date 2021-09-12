package main

import "fmt"

func main() {
	fmt.Println(IterativePower(4, 3))
}

func IterativePower(nb int, power int) int {
	if nb < 0 || nb > 25 {
		return 0
	} else if nb == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= power; i++ {
		result = result * i
	}
	return result
}
