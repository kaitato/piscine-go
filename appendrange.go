package piscine

func AppendRange(min, max int) []int {
	var result []int
	for i := min; i < max-1; i++ {
		if min > 0 {
			result = append(result, i+1)
		} else {
			return result
		}
	}
	return result
}
