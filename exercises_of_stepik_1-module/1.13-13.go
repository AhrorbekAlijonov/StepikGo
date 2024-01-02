package main

import "fmt"

func main() {

	var first, second, third, iter, a int
	first = 1
	second = 1
	iter = 2
	fmt.Scan(&a)
	for second < a {
		iter++
		third = second + first
		first = second
		second = third
	}
	if third == a {
		fmt.Println(iter)
	} else {
		fmt.Println(-1)
	}

}
