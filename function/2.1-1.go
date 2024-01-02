package main

import "fmt"

func main() {
	var str string
	fmt.Scanln(&str)
	f(str)
}

func f(text string){
        fmt.Println(text)
}