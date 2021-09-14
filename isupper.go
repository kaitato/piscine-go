package piscine

func IsUpper(s string) bool {

	count := 0
	for _, caps := range s {
		if caps >= 'A' && caps <= 'Z' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
