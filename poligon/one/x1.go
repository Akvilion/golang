package main

import "fmt"

func adder(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(adder(1, 3))
}
