package main

import (
	"fmt"
	"time"
)

func main() {
    var tm string
	fmt.Scanln(&tm)

	unx, err := time.Parse(time.RFC3339, tm)
	if err != nil {
		panic(err)
	}
	unix := unx.Format("Mon Jan _2 15:04:05 MST 2006")
	fmt.Println(unix)
}