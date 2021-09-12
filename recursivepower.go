package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if nb > 1 {
		return nb * RecursivePower(nb-1, power)
	}
	return 0
}
