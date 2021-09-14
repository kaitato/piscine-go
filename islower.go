package piscine

func IsLower(s string) bool {

	count := 0
	for _, caps := range s {
		if caps >= 'a' && caps <= 'z' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
