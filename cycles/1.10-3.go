package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)
	sum := 0
	var z int
	for i := 1; i <= a; i++ {
		fmt.Scan(&z)
		if  9 < z && z < 100 && z % 8 == 0 {
			sum += z
		}
	}
	fmt.Println(sum)
}