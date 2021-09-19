package piscine

func SplitWhiteSpaces(s string) []string {
	var aString []string
	count := 1
	strlen := 0
	result := ""
	for c := range s {
		if isWhiteSpace(s[c]) {
			count++
		}
		strlen++
	}
	aString = make([]string, count)
	i := 0
	for j, c := range s {
		if j+1 == strlen {
			aString[i] = result + string(s[j])
		}
		if isWhiteSpace(s[j]) {
			if i <= count {
				aString[i] = result
				i++
			}
		} else {
			result += string(c)
		}
	}
	return aString
}

func isWhiteSpace(r byte) bool {
	if r == ' ' || r == '\n' || r == '\t' {
		return true
	}
	return false
}
