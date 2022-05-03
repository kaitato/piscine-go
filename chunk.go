package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Chunk(a []int, ch int) {
	var slice []int
	if ch <= 0 {
		// fmt.Println()
		z01.PrintRune('\n')
		return
	}
	result := make([][]int, 0, len(a)/ch+1)
	for len(a) >= ch {
		slice, a = a[:ch], a[ch:]
		result = append(result, slice)
	}
	if len(a) > 0 {
		result = append(result, a[:])
	}
	fmt.Println(result)
}
