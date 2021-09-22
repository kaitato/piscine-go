package main

import (
	"os"

	"github.com/01-edu/z01"
)

// print digits
func main() {
	for i := '0'; i <= '9'; i++ { // write a loop initialing i as 0 to 9
		z01.PrintRune(i) // print i
	}
	z01.PrintRune('\n') // print a new line

	// display alpha m
	z01.PrintRune('m')
	z01.PrintRune('\n')

	// display rev m ? maybe same as display m
	z01.PrintRune('m')
	z01.PrintRune('\n')

	// countdown
	for a := '9'; a >= '0'; a-- { // loop for first digit
		for b := '9'; b >= '0'; b-- { // loop for second digit
			for c := '9'; c >= '0'; c-- { // loop for third digit
				z01.PrintRune(a) // print each number
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a == '0' && b == '0' && c == '0' { // if digits hit the last number
					z01.PrintRune('\n') // print new line
				} else { // otherwise print ', '
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}

// firstrune
func FirstRune(s string) rune {
	a := []rune(s)
	return a[0]
}

// lastrune
func LastRune(s string) rune {
	a := []rune(s)
	return a[len(a)-1]
}

// paramcount
func main() {
	return len(os.Args)
}

// rot14
func Rot14(s string) string {
	a := []rune(s)
	var result string
	for i := 0; i < len(a); i++ {
		if a[i] >= 'a' && a[i] <= 'l' || a[i] >= 'A' && a[i] <= 'L' {
			a[i] += 14
		} else if a[i] >= 'm' && a[i] <= 'z' || a[i] >= 'M' && a[i] <= 'Z' {
			a[i] -= 12
		}
		result += string(a[i])
	}
	return result
}

// max
func Max(a []int) int {
	max := 0
	for i := range a {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

// wdmatch

// is power of 2

// search replace

// compare
func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// doop

// expand str

// reverse bits

// atoi

// print bits

// reverse str cap

// tabmult

// capitalize
func Capitalize(s string) string {
	aRune := []rune(s)
	for i := range aRune {
		if aRune[i] >= 'A' && aRune[i] <= 'Z' {
			aRune[i] = aRune[i] + 32
		}
	}

	for i, char := range aRune {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			if i == 0 || aRune[i-1] >= 32 && aRune[i-1] <= 47 || aRune[i-1] >= 58 && aRune[i-1] <= 64 || aRune[i-1] >= 91 && aRune[i-1] <= 96 || aRune[i-1] >= 122 {
				if aRune[i] >= 'a' && aRune[i] <= 'z' {
					aRune[i] = char - 32
				}
			} else {
				if aRune[i] >= 'A' && aRune[i] <= 'Z' {
					aRune[i] = char + 32
				}
			}
		}
	}
	return string(aRune)
}

// gcd

// add prime sum

// clean str

// reverse hex

// reverse range

// hidden p

// range

// f prime

// for each
func ForEach(f func(int), a []int) {
	for _, a := range a {
		f(a)
	}
}

// atoi base

// itao

// brainfuck

// itaobase
