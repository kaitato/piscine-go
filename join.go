package piscine

func Join(strs []string, sep string) string {
	result := ""
	for index, letters := range strs {
		result += letters
		if strs[index] == strs[len(strs)-1] {
			break
		} else {
			result += sep
		}
	}
	return result
}
