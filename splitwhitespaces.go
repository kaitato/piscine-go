package piscine

func SplitWhiteSpaces(s string) []string {
	ln := 0
	spaces := false
	for c := range s {
		if spaces && c != 0 && (s[c-1] == '\n' ||
			s[c-1] == '\t' || s[c-1] == ' ') &&
			s[c] != '\n' && s[c] != '\t' && s[c] != ' ' {
			ln++
		}
		if s[c] != '\n' && s[c] != '\t' && s[c] != ' ' {
			spaces = true
		}
	}
	ln++

	count := 0
	result := make([]string, ln)
	ok := true
	for _, c := range s {
		if c == '\n' || c == '\t' || c == ' ' {
			if !ok {
				count++
			}
			ok = true
			continue
		}
		result[count] = result[count] + string(c)
		ok = false
	}
	return result

}
