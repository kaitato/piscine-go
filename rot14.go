package piscine

func rotate(b rune) rune {
	if b >= 'A' && b < 'M' || b >= 'a' && b < 'm' {
		return b + 14
	}
	if b >= 'M' && b <= 'Z' || b >= 'm' && b <= 'z' {
		return b - 12
	}
	return b
}

func Rot14(s string) string {
	result := ""
	for _, char := range s {
		result += string(rotate(char))
	}
	return result
}
