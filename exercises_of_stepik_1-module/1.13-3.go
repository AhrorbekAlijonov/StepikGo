package main

import "fmt"

func main() {
	var a int
	b := 0
	fmt.Scan(&a)
	b = a / 3600
	a = a % 3600 / 60
	fmt.Printf("It is %d hours %d minutes.", b, a)
}
