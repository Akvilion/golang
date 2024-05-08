package main

import "fmt"

func one() {

	if r := recover(); r != nil {
		fmt.Println("recover block", r) // recover block panic
		// panic("33333")
	}
}

func two() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover block2222", r)
			panic("another panic2")
			// recover block panic
		}
	}()
	panic("another panic1")

}

func main() {
	defer one()

	fmt.Println("1")
	// panic("panic")
	two()
	fmt.Println("2")
}
