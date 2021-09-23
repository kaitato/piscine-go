package piscine

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	for a := 0; a <= len(arg); a++ {
		for i := 1; i < len(arg); i++ {
			if arg[i] < arg[i-1] {
				arg[i], arg[i-1] = arg[i-1], arg[i]
			}
		}
	}
	return arg[2]
}
