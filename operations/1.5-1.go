package main

import "fmt"

func main() {
	var a int 
	fmt.Scanln(&a)
	a = a * 2 
	a = a + 100 
	fmt.Println(a)
}