package main

import "fmt"

func one() *int { // повертаєме значення типу вказівник
	a := 2
	return &a // повертаємо вказівник
}

func main() {
	// var myInt int = 42
	// var myIntPointer *int = &myInt
	// fmt.Println(*myIntPointer)

	f := one()
	fmt.Println(*f)

}
