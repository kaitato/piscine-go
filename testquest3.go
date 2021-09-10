package piscine

import ""

//take a pointer to an int as argument and give thi int the value of 1
func PointOne(n *int) {
	*n = 1
}

//take a pointer to a pointer to a pointer to an int as argument and give this in the value of 1
func UltimatePointOne(n ***int) {
	a := 1
	*b := &a
	**c := &b
	***n = **c
}

//divide int a and b store as int pointed by div and mod
func DivMod(a int,b int ,div *int,mod *int,) {
	*div = a / b
	*mod = a % b
}

// divide int a and b and store div pointed by a and mod pointed by b
func UltimateDivMod(a *int, b *int) {
	c := *a / *b
	d := *a % *b

	*a = c
	*b = d
}

//Prints string of characters
func PrintStr(s string) {
	for _, a := range s {
		z01.PrintRune(a)
	}
}

//counts the runes of a string and returns count
func StrLen(s string) int {
	return len(s)
}

//take two pointers and swap contents
func Swap(a *int, b * int) {
	*a, *b = *b, *a
}

//print reverse string
func StrRev(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i<j;i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}