package main

import "github.com/01-edu/z01"

func main() {
	// printalphabet
	for a := 'a'; a <= 'z'; a++ {
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')

	// print alphabet reverse
	for b := 'z'; b >= 'a'; b-- {
		z01.PrintRune(b)
	}
	z01.PrintRune('\n')

	// print digits
	for c := '0'; c <= '9'; c++ {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')

	// is negative
	d := -5
	if d >= 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
	z01.PrintRune('\n')

	// print comb
	for e := '0'; e <= '9'; e++ {
		for f := '0'; f <= '9'; f++ {
			for g := '0'; g <= '9'; g++ {
				if e < f && f < g {

					z01.PrintRune(e)
					z01.PrintRune(f)
					z01.PrintRune(g)
					if e < '7' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('\n')
					}
				}
			}
		}
	}
}
