package piscine

func TrimAtoi(s string) int {
	neg := false
	str := []rune(s)
	result := 0
	for i := 0; i < len(str); i++ {
		if !neg && result == 0 && str[i] == '-' {
			neg = true
		}
		if str[i] >= '0' && str[i] <= '9' {
			result *= 10
			result += int(str[i] - 48)
		}
	}
	if neg {
		return result * -1
	}
	return result
}
