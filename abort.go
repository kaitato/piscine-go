package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	sort.Ints(arg)
	return arg[2]
}
