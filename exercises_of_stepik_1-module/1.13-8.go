package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	var array []int
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		array = append(array, b)
	}
	count := 0
	min_num := array[0]
	for i := 0; i < a; i++ {
		if min_num < array[i] {
			min_num = array[i]
			count = 1
		} else if min_num == array[i] {
			count++
		}
	}
	fmt.Println(count)
}
