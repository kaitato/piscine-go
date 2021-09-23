package main

import (
	"fmt"
	"os"
	"strconv"
)

func ChechOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" || s == "%" {
		return true
	}
	return false
}

func main() {
	result := 0
	num1 := os.Args[1]
	operator := os.Args[2]
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
	if ChechOperator(operator) == true {
		if operator == "+" {
			result = nb1 + nb2
		} else if operator == "-" {
			result = nb1 - nb2
		} else if operator == "*" {
			result = nb1 * nb2
		} else if operator == "/" {
			if nb2 == 0 {
				fmt.Println("No division by 0")
				return
			}
			result = nb1 / nb2
		} else if operator == "%" {
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
