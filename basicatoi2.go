package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, number := range s {
		if number >= 0 && number <= 9 {
			conv := int(number) - 48
			result = (result * 10) + conv
		} else {
			return 0
		}
	}
	return result
}
