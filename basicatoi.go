package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, number := range s {
		conv := int(number) - 48
		result = result*10 + conv
	}
	return result
}
