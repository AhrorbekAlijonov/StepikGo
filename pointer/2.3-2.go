package main

import "fmt"

func main()  {
	var x, y int
	fmt.Scan(&x, &y)
	test(&x, &y)
}

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
 }