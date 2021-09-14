package piscine

func ToLower(s string) string {
	aRune := []rune(s)
	for i := range aRune {
		if aRune[i] >= 'A' && aRune[i] <= 'Z' {
			aRune[i] = aRune[i] + 32
		}
	}
	return string(aRune)
}
