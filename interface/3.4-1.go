package main

import (
  "fmt"           
)

func readTask() (interface, interface, interface) {
	var a, b, c interface
	fmt.Scan(&a, &b, &c)
	return a, b, c
}

func main() {
  a, b, c := readTask()

  if _, ok := a.(float64); !ok {
    fmt.Printf("value=%v: %T", a, a)
    return
  }

  if _, ok := b.(float64); !ok {
    fmt.Printf("value=%v: %T", b, b)
    return
  }

  if val, ok := c.(string); ok {
    switch val {
    case "+":
            fmt.Printf("%.4f", a.(float64) + b.(float64))
    case "-":
            fmt.Printf("%.4f", a.(float64)- b.(float64))
    case "*":
            fmt.Printf("%.4f", a.(float64) * b.(float64))
    case "/":
            fmt.Printf("%.4f", a.(float64) / b.(float64))
    default:
      fmt.Println("неизвестная операция")
    }
  } else {
    fmt.Println("неизвестная операция")
  }
}
