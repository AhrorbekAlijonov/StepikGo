package main

import (
	"fmt"
	"strings"
)

func findCommonDigits(num1, num2 int) {
	// intni stringga o'girish
	num1Str := fmt.Sprint(num1)
	num2Str := fmt.Sprint(num2)

	result := ""

	for _, digit1 := range num1Str {
		for _, digit2 := range num2Str {
			if digit1 == digit2 && strings.IndexRune(result, digit1) == -1 {
				result += string(digit1) + " "
			}
		}
	}

	fmt.Println(result)
}

func main() {
	var num1, num2 int
	fmt.Scan(&num1)

	fmt.Scan(&num2)

	findCommonDigits(num1, num2)
}
