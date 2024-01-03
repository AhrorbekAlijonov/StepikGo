package main

import (
	"bufio"
	"strconv"
	"os" 
	"io"
)

func main() {
    buf := bufio.NewScanner(os.Stdin)
  var (
    sum int
  )
  for buf.Scan() {
    num, err := strconv.Atoi(buf.Text())
    if err != nil {
      return
    }
    sum += num
  }
  io.WriteString(os.Stdout, strconv.Itoa(sum))
}
