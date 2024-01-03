package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	tm, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	tim, _ := time.Parse("2006-01-02 15:04:05", tm[:19])
	ti := tim.Format("2006-01-02 15:04:05")

	if tim.Hour() < 13 {
		fmt.Println(ti)
        return
	} else {
		tim = tim.AddDate(0, 0, 1)
	}

	t := tim.Format("2006-01-02 15:04:05")
	fmt.Println(t)
}
