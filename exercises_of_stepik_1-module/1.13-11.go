package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%10 == 1 && n != 11 {
		fmt.Printf("%d korovy", n)
	} else if n%10 > 1 && n%10 < 5 && n/10 != 1 {
		fmt.Printf("%d korov", n)
	} else {
		fmt.Printf("%d korova", n)
	}
}
