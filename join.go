package piscine

import "github.com/01-edu/z01"

func Join(strs []string, sep string) string {
	result := ""
	for index, letters := range strs {
		result += letters
		if strs[index] == strs[len(strs)-1] {
			z01.PrintRune('\n')
		} else {
			result += sep
		}
	}
	return result
}
