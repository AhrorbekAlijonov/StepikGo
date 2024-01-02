package main

import (
	"fmt"
	"math"
)

func main() {
	// Ввод числа типа float64
	var input float64
	fmt.Scan(&input)

	// Проверка условий и форматирование вывода
	if input <= 0 {
		fmt.Printf("число %.2f не подходит\n", input)
	} else if input > 10000 {
		fmt.Printf("%e\n", input)
	} else {
		result := math.Pow(input, 2)
		// Округление до 4 знаков после запятой
		result = math.Round(result*10000) / 10000
		fmt.Printf("%.4f\n", result)
	}
}
