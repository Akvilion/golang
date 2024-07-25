package main

import "fmt"

func one() {
	// recover()
	fmt.Println(recover()) // в цьому випадку повернеться те що в panic
}

func two() {
	defer one() // оце заглушило паніку
	panic("oh no")
	fmt.Println("I won't be run!")
}

func main() {
	two()
	fmt.Println("Exiting normally")
}
