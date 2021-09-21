package piscine

func ForEach(f func(int), a []int) {
	for _, a := range a {
		f(a)
	}
}
