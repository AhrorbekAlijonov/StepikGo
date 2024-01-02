package main

import (
	"fmt"
)

func main() {
	var a , b string
	fmt.Scan(&a)
	for _,v := range a {
		b = string(v) + b
	  }

	if a == b {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}