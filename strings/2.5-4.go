package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	for i, _ := range a {
		if i%2 != 0 {
            fmt.Print(string(a[i]))
		}
	}
}