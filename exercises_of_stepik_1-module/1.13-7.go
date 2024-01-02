package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	count := 0
	for i := 0; i < a; i++ {
		var zero int
		fmt.Scan(&zero)
		if zero == 0 {
			count++
		}
	}
	fmt.Println(count)
}
