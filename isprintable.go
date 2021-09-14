package piscine

func IsPrintable(s string) bool {
	length := len(s)
	count := 0
	for _, caps := range s {
		if caps > 32 {
			count++
		}
	}
	if count == length {
		return false
	}
	return true
}
