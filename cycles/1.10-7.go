package main

import "fmt"

func main() {
	a, b, c, i := 0, 0, 0, 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	for {
		a = a + a*b/100
		i++
		if a >= c {
			break
		}
	}
	fmt.Println(i)
}