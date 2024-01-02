package main

import "fmt"

func main() {
    var z int

    fmt.Scan(&z)

    a := z % 10
    b := z / 10 % 10
    c := z / 100
    if a == b {
        fmt.Println("NO")
    } else if b == c {
        fmt.Println("NO")
    } else if a == c {
        fmt.Println("NO")
    } else {
        fmt.Println("YES")
    }
    }