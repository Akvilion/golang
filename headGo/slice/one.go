package main

import "fmt"

func main() {
	var a []int

	fmt.Println(a) // []

	b := make([]int, 2)
	// b[2] = 2
	fmt.Println(b) // [0 0 0 0 0 0 0]

}
