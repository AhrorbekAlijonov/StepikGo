package main

import (
	"fmt"
	"strconv"
	"strings"
)

func removeDigit(number int, digitToRemove int) int {
	numberStr := strconv.Itoa(number)

	resultStr := strings.ReplaceAll(numberStr, strconv.Itoa(digitToRemove), "")

	result, err := strconv.Atoi(resultStr)
	if err != nil {
		fmt.Println("Xato sodir bo'ldi:", err)
		return number
	}

	return result
}

func main() {
	var number, digitToRemove int
	fmt.Scan(&number)

	fmt.Scan(&digitToRemove)

	result := removeDigit(number, digitToRemove)
	fmt.Println(result)
}