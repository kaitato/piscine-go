package piscine

func BasicAtoi(s string) int {
	j := 0
	for _, i := range s {
		a := int(i) - 48
		j = j*10 + a
	}
	return j
}
