package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	s := make([]int, a, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Println(s[3])
}