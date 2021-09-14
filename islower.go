package piscine

func IsLower(s string) bool {
	length := len(s)
	count := 0
	for _, caps := range s {
		if caps >= 'a' && caps <= 'z' {
			count++
		}
	}
	if count == length {
		return true
	}
	return false
}
