package main

import "fmt"

func main() {
	var x, y int
	var a, b bool
	fmt.Scan(&x, &y)
	if y%x == 0 {
		a = true
	}
	if x%y == 0 {
		b = true
	}
	fmt.Print(a, "\n", b)
}
