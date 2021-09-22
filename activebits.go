package piscine

func ActiveBits(n int) int {
	count := 0
	if n < 1 || n > 64 {
		return 0
	}
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}
	return count
}
