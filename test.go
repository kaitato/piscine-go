package main

import "github.com/01-edu/z01"

func main() {
	QuadA(5, 3)
}

func QuadA(x, y int) {
	if x >= 1 || y >= 1 {
		z01.PrintRune('o')
		for i := 2; i < x; i++ {
			z01.PrintRune('-')
		}
		if x > 1 {
			z01.PrintRune('o')
		}
		z01.PrintRune('\n')
	}

	if y > 2 {
		for i := 2; i < y; i++ {
			z01.PrintRune('|')
			for j := 2; j < x; j++ {
				z01.PrintRune(' ')
			}
			if x > 1 {
				z01.PrintRune('|')
				z01.PrintRune('\n')
			}

		}
	}
	if y > 1 {
		z01.PrintRune('o')
		for i := 2; i < x; i++ {
			z01.PrintRune('-')
		}
		if x > 1 {
			z01.PrintRune('o')
		}
		z01.PrintRune('\n')
	}
}
