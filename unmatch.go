package piscine

func Unmatch(a []int) int {
	for _, i := range a {
		count := 0
		for _, j := range a {
			if i == j {
				count++
			}
		}
		if count == 1 {
			return i
		}
	}
	return -1
}
