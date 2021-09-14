package piscine

func IsPrintable(s string) bool {
	count := 0
	for _, caps := range s {
		if caps >= 32 && caps <= 127 {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
