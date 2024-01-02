package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	max := 0
	for i := 0; i < 5; i++ {
		if array[i] > max {
			fmt.Println((array[i]))
			max = array[i]
		}
	}
	fmt.Println(max)
}
