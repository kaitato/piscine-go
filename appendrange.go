package piscine

func AppendRange(min, max int) []int {
	var result []int
	for i := min; i <= max-1; i++ {
		result = append(result, i+1)
	}
	return result
}