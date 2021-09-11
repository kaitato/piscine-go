package piscine

func BasicAtoi(s string) int {
	a := []rune(s)
	for _, i := range a {
		for j := '0'; j < i; j++ {
			return a[j]
		}
	}
}
