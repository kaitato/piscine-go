package piscine

func IsUpper(s string) bool {
	length := len(s)
	count := 0
	for _, caps := range s {
		if caps >= 'A' && caps <= 'Z' {
			count++
		}
	}
	if count == length {
		return true
	}
	return false
}
