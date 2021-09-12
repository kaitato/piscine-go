package piscine

func IterativePower(nb int, power int) int {
	a := nb
	if nb < 0 {
		return 0
	}
	for i := 1; i < power; i++ {
		a *= nb
	}
	return a
}
