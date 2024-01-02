package main

import "fmt"

func main() {
    var a int

    fmt.Scan(&a)

    b := a / 1000
    
    a = a % 1000

    sum := (b % 10) + (b / 10 % 10) + (b / 100)
    sum2 := (a % 10) + (a / 10 % 10) + (a / 100)

    if sum == sum2 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
    }