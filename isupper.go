package piscine

func IsUpper(s string) bool {
	aRune := len(s)
	count := 0
	for _, caps := range s {
		if caps >= 'A' && caps <= 'Z' {
			count++
		}
	}
	if count == aRune {
		return true
	}
	return false
}
