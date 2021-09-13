package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	b := 0
	for i := range s {
		b = i + 1
	}
	if n <= b && n-1 >= 0 {
		return a[n-1]
	} else {
		return 0
	}
}
