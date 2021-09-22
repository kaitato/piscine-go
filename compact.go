package piscine

func Compact(ptr *[]string) int {
	var s []string
	for _, i := range *ptr {
		if i != "" {
			s = append(s, i)
		}
		*ptr = s
	}
	return len(s)
}
