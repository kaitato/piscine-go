package piscine

func BasicJoin(elems []string) string {
	result := ""
	for _, elems := range elems {
		result += elems
	}
	return result
}
