package main

import "fmt"

func main() {
	var a, b , c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(vote(a, b, c))
}

func vote(x int, y int, z int) int {
    if x == y {
        return x
    }else if y == z {
        return y
    }else {
        return z
    }
}