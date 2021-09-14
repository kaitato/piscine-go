package piscine

func Index(s string, toFind string) int {

	a := -1
	b := []rune(toFind)
	for k, c := range s {
		if k == StrLen(s) {
			break
		}
		if c == b[0] {
			a = k
		} else {
			continue
		}
	}
	return a
}

func StrLen(str string) int {
	count := 0
	for _, _ = range str {
		count++
	}
	return count
}
