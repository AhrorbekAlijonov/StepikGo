package main

import (
	"fmt"
	"strings"
)
func main() {
	var a string
	fmt.Scan(&a)
	for _, el := range a {
		count := strings.Count(a, string(el))
		if count == 1 {
			fmt.Print(string(el))
		}
	}
}