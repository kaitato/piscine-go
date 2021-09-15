package main

import "fmt"

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

func TrimAtoi(s string) int {
	neg := false
	str := []rune(s)
	result := 0
	for i := 0; i < len(str); i++ {
		if !neg && result == 0 && str[i] == '-' {
			neg = true
		}
		if str[i] >= '0' && str[i] <= '9' {
			result *= 10
			result += int(str[i] - 48)
		}
	}
	if neg {
		return result * -1
	}
	return result
}
