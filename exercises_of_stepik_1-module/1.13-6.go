package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scanln(&a, &b)
	a = (a + b) / 2
	fmt.Printf("%v", a)
}
