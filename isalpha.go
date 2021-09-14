package piscine

func IsAlpha(s string) bool {
	length := len(s)
	count := 0
	for _, caps := range s {
		if caps >= 'a' && caps <= 'z' || caps >= 'A' && caps <= 'Z' || caps >= '0' && caps <= '9' {
			count++
		}
	}
	if count == length {
		return true
	}
	return false
}
