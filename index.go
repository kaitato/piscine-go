package main

import "fmt"

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

func StrLen(str string) int {
	count := 0
	for _, _ = range str {
		count++
	}
	return count
}

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
