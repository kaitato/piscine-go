package piscine

func IsLower(s string) bool {
	aRune := len(s)
	count := 0
	for _, caps := range s {
		if caps >= 'a' && caps <= 'z' {
			count++
		}
	}
	if count == aRune {
		return true
	}
	return false
}
