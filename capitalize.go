package piscine

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
