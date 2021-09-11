package piscine

func Atoi(s string) int {
	result := 0
	for _, number := range s {
        if number == '-' {
            return number
        }
		if number >= 48 && number <= 57 {
			conv := int(number) - 48
			result = (result * 10) + conv
		} else {
			return 0
		}
	}
	return result
}