package piscine

func Capitalize(s string) string {
	aRune := []rune(s)
	for i, char := range aRune {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			if i == 0 || aRune[i-1] >= 'a' && aRune[i-1] <= 'z' || aRune[i-1] >= 'A' && aRune[i-1] <= 'Z' || aRune[i-1] >= '0' && aRune[i-1] <= '9' {
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
