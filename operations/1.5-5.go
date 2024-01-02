package main    

import "fmt"

func main() {
    var i int
    fmt.Scanln(&i)
    fmt.Println(i % 100 / 10)
}