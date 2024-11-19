package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x < 0 && x%2 == 0 {
		fmt.Print("genap negatif")
	} else {
		fmt.Print("bukan")
	}
}
