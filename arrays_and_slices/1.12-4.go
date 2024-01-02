package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	arr := make([]int, a, a)
	var array []int
	for i := 0; i < a; i++ {
		var b int
		fmt.Scan(&b)
		arr[i] = b
	}
	for i := 0; i < a; i++ {
		if i%2 == 0 {
			array = append(array, arr[i])
		}
	}
    for i := 0; i < len(array); i++ {

        fmt.Print(array[i], " ")
    }
}
