package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var result []int = nil
		return result
	} else {
		result := make([]int, max-min)
		j := 0
		for i := min; i < max; i++ {
			result[j] = i
			j++
		}
		return result
	}
}
