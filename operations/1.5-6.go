package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    fmt.Println("It is", a / 30, "hours", a % 30 *2, "minutes.")
}