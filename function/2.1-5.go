package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(sumInt(a, b, c, d))
}

func sumInt(nums ...int) (int, int) {
	sum := 0
	count := 0
	for _, a := range nums {
		sum += a
		count++
	}
	return count, sum
}