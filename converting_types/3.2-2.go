package main

import (
	"fmt"
	"unicode"
	"strconv"
)

func main() {
	var a , b string
	fmt.Scan(&a, &b)
	fmt.Println(adding(a, b))
}

func adding(a, b string) int64 {
	c := []rune(a)
	var s, m string
	for i, _ := range c {
		if unicode.IsDigit(c[i]) {
			s += string(c[i])
		}
	}

	d := []rune(b)
	for i, _ := range d {
		if unicode.IsDigit(d[i]) {
			m += string(d[i])
		}
	}

	sInt, _ := strconv.Atoi(s)
	mInt, _ := strconv.Atoi(m)

	result := sInt + mInt
    var res int64
    res = int64(result)
	return res
}