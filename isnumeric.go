package piscine

func IsNumeric(s string) bool {
	count := 0
	for _, caps := range s {
		if caps >= '0' && caps <= '9' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
