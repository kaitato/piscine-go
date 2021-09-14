package piscine

func IsNumeric(s string) bool {
	length := len(s)
	count := 0
	for _, caps := range s {
		if caps >= '0' && caps <= '9' {
			count++
		}
	}
	if count == length {
		return true
	}
	return false
}
