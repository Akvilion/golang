package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting the program")

	// This will cause a panic
	causePanic()

	fmt.Println("This line will not be reached")
}

func causePanic() {
	fmt.Println("About to panic")
	panic("Something went wrong!")
}
