package piscine

func ToUpper(s string) string {
	aRune := []rune(s)
	for i := range aRune {
		if aRune[i] >= 'a' && aRune[i] <= 'z' {
			aRune[i] = aRune[i] - 32
		}
	}
	return string(aRune)
}
