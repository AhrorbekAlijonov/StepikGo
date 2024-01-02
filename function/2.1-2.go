package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}



func minimumFromFour() int {
    //print your code here
    min_num := 99999999999
	var b int
	for i := 0; i < 4; i++ {

		fmt.Scan(&b)
		if min_num > b {
			min_num = b
		}
	}
	return min_num
}