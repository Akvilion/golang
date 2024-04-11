package main

import "fmt"

func main() {
	var myInt int = 42
	var myIntPointer *int = &myInt
	fmt.Println(*myIntPointer)
}
