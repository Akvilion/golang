package main

import "fmt"

func one() {
	recover()
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
