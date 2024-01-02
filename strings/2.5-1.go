package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text , _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// text = rune(text)
	var run = []rune(text)

	if unicode.IsUpper(run[0]) && string(run[len(run)-1]) == "." {
		fmt.Println("Right")

	} else {
		fmt.Println("Wrong")
	}
}