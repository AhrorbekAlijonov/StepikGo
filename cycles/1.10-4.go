package main

import "fmt"

func main() {
	count := 0
	max := 0

	for {
		var a int

		fmt.Scanln(&a)
		if a == 0 {
			break
		}
		if a > max {
			max = a
			count = 1
		} else if a == max {
			count++
		}
	}
	fmt.Println(count)
}