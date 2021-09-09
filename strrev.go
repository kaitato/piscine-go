package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string {
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune [n] = range
		n++
	}
	rune = rune [0:n]
	for i := 0; i < n / 2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	output := string(rune)
}
