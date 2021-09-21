package piscine

func Map(f func(int) bool, a []int) []bool {
	mp := make([]bool, len(a))

	for i, j := range a {
		mp[i] = f(j)
	}
	return mp
}
