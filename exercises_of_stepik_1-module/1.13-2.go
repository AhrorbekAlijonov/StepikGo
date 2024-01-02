package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	for a%10 != 0 {
		b += a % 10
		b *= 10
		a /= 10
	}
	fmt.Println(b / 10)
}
