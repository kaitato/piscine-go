package piscine

func LastRune(s string) rune {
	aRune := []rune(s)
	return aRune[len(s)-1]
}
