package main

import "fmt"

func one() *int {
	a := 2
	return &a
}

func main() {
	// var myInt int = 42
	// var myIntPointer *int = &myInt
	// fmt.Println(*myIntPointer)

	f := one()
	fmt.Println(*f)

}
