package piscine

func ToUpper(s string) string {
	Arune := []rune(s)
	for i := range Arune {
		if Arune[i] >= 'a' && aRune <= 'z' {
			aRune[i] = aRune[i] - 32
		}
	}
	return string(aRune)
}
