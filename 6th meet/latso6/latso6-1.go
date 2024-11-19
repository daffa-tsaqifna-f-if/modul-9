package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	if x == 1 {
		fmt.Print(1)
	} else {
		y = x / 2
		x = x % 2
		if x == 1 {
			fmt.Print(y + 1)
		} else {
			fmt.Print(y)
		}
	}
}
