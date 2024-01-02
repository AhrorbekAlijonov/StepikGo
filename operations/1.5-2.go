package main

import "fmt"

func main(){

	var a int
	var b int
	fmt.Scan(&a) // считаем переменную 'a' с консо
	fmt.Scan(&b) // считаем переменную 'b' с консоли
  
	a = a * a
	b = b * b
	a = a+b
	fmt.Println(a)
  }