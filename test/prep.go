package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

// print digits
func main() {
	for i := '0'; i <= '9'; i++ { // write a loop initialing i as 0 to 9
		z01.PrintRune(i) // print i
	}
	z01.PrintRune('\n') // print a new line

	// display alpha m
	for lower, upper := 'a', 'B'; upper <= 'Z'; lower, upper = lower+2, upper+2 {
		z01.PrintRune(lower)
		z01.PrintRune(upper)
	}

	// display rev m
	for lower, upper := 'z', 'Y'; upper >= 'A'; lower, upper = lower-2, upper-2 {
		z01.PrintRune(lower)
		z01.PrintRune(upper)
	}

	// countdown
	for a := '9'; a >= '0'; a-- { // loop for first digit
		for b := '9'; b >= '0'; b-- { // loop for second digit
			for c := '9'; c >= '0'; c-- { // loop for third digit
				z01.PrintRune(a) // print each number
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a == '0' && b == '0' && c == '0' { // if digits hit the last number
					z01.PrintRune('\n') // print new line
				} else { // otherwise print ', '
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}

// firstrune
func FirstRune(s string) rune {
	a := []rune(s)
	return a[0]
}

// lastrune
func LastRune(s string) rune {
	a := []rune(s)
	return a[len(a)-1]
}

// paramcount
func main() {
	fmt.Println(len(os.Args) - 1) // if can use fmt

	digit1 := 48
	for a := 0; a < len(os.Args); a++ {
		digit1 += 1
	}
	toreturn := rune(digit1) // if count <= 9
	z01.PrintRune(toreturn)
	z01.PrintRune('\n')
}

// rot14
func Rot14(s string) string {
	a := []rune(s)
	result := ""
	for i := 0; i < len(a); i++ {
		if a[i] >= 'a' && a[i] <= 'l' || a[i] >= 'A' && a[i] <= 'L' {
			a[i] += 14
		} else if a[i] >= 'm' && a[i] <= 'z' || a[i] >= 'M' && a[i] <= 'Z' {
			a[i] -= 12
		}
		result += string(a[i])
	}
	return result
}

// max
func Max(a []int) int {
	max := 0
	for i := range a {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

// wdmatch
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		string1 := []rune(os.Args[1])
		string2 := []rune(os.Args[2])
		matchFound := 0
		secondStringStart := 0

		if len(string1) < len(string2) {
			for i := 0; i < len(string1); i++ {
				for j := secondStringStart; j < len(string2); j++ {
					if string(string1[i]) == string(string2[j]) {
						secondStringStart = j
						matchFound = matchFound + 1
						break
					}
				}
			}
			if matchFound == len(string1) {
				fmt.Println(string(string1))
			}
		}
	}
}




// is power of 2
func ispowerof2(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb >= 2 {
		result := 1
		for i := 1; result <= nb; i++ {
			result = result * 2
			if result == nb {
				return true
			}
		}
		return false
	} else {
		return false
	}
}

// search replace
func main() {
	var result []rune // fmt
	arg := os.Args
	search := []rune(arg[2])
	replace := []rune(arg[3])
	if len(arg) != 4 {
		return
	}
	for _, i := range arg[1] {
		if search[0] == i {
			result = append(result, replace[0]) // fmt
			z01.PrintRune(replace[0])
		} else {
			result = append(result, i) // fmt
			z01.PrintRune(i)
		}
	}
	fmt.Printf(string(result)) // fmt
	z01.PrintRune('\n')
}




// compare
func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// doop
func CheckOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" || s == "%" {
		return true
	}
	return false
}

func main() {
	result := 0
	num1 := os.Args[1]
	op := os.Args[2]
	num2 := os.Args[3]
	nb1, err1 := strconv.Atoi(num1)
	if err1 != nil {
		fmt.Println("0")
		return
	}
	nb2, err2 := strconv.Atoi(num2)
	if err2 != nil {
		fmt.Println("0")
		return
	}
	if CheckOperator(op) == true {
		if op == "+" {
			result = nb1 + nb2
		} else if op == "-" {
			result = nb1 - nb2
		} else if op == "*" {
			result = nb1 * nb2
		} else if op == "/" {
			if nb2 == 0 {
				fmt.Println("No division by 0")
				return
			}
			result = nb1 / nb2
		} else if op == "%" {
			if nb2 == 0 {
				fmt.Println("No Modulo by 0")
				return
			}
			result = nb1 % nb2
		}
		fmt.Println(result)
	} else {
		fmt.Println("0")
		return
	}
}



// expand str

// reverse bits

// atoi
func Atoi(str string) int {
	result := 0
	val := 1
	for i, num := range str {
		conv := int(num) - 48
		if conv <= 9 && conv >= 0 {
			result = (result * 10) + conv
		} else if conv == -3 && i == 0 {
			val = -1
		} else if conv == -5 && i == 0 {
			val = 1
		} else {
			return 0
		}
	}
	result *= val
	return result
}


// print bits

// reverse str cap

// tabmult

// capitalize
func Capitalize(s string) string {
	aRune := []rune(s)
	for i := range aRune {
		if aRune[i] >= 'A' && aRune[i] <= 'Z' {
			aRune[i] = aRune[i] + 32
		}
	}

	for i, char := range aRune {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			if i == 0 || aRune[i-1] < '0' || aRune[i-1] > '9' && aRune[i-1] < 'A' || aRune[i-1] > 'Z' && aRune[i-1] < 'a' || aRune[i-1] > 'z' {
				if aRune[i] >= 'a' && aRune[i] <= 'z' {
					aRune[i] = char - 32
				}
			} else {
				if aRune[i] >= 'A' && aRune[i] <= 'Z' {
					aRune[i] = char + 32
				}
			}
		}
	}
	return string(aRune)
}

// gcd

// add prime sum

// clean str

// reverse hex

// reverse range

// hidden p

// range
func main() {
	argu := os.Args
	r1, _ := strconv.Atoi(argu[1])
	r2, _ := strconv.Atoi(argu[2])
	if len(argu) != 3 {
		fmt.Print()
	} else {
		if r1 < r2 {
			for i := r1; i <= r2; i++ {
				fmt.Printf("%v ", i)
			}
		} else if r1 > r2 {
			for j := r1; j >= r2; j-- {
				fmt.Printf("%v ", j)
			}
		}
	}
}




// f prime

// for each
func ForEach(f func(int), a []int) {
	for _, a := range a {
		f(a)
	}
}

// atoi base

// itao

// brainfuck

// itaobase
