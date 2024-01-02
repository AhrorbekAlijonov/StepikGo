package main

import "fmt"

func main()  {
	var a int64
	fmt.Scan(&a)
	fmt.Println(convert(a))
}

func convert(a int64) uint16{
	var b uint16
	b = uint16(a)
	return b
}